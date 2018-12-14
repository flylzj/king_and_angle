package chat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"model"
	"net/http"
	"resource"
)
var clients = make(map[string]*websocket.Conn)
var broadcast = make(chan model.ChatMessage)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsConnectionHandle(ctx *gin.Context){
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil{
		fmt.Println("connect err", err.Error())
	}

	defer ws.Close()

	var msg model.ChatMessage

	if err := ws.ReadJSON(&msg);err != nil{
		ws.WriteJSON(model.NoticeMessage{Message:"data error", Code:1})
	}

	claims, result := resource.CheckToken(msg.Token)
	if result{
		currentUser := resource.GetUserById(claims.ID)
		clients[currentUser.Username] = ws

		for{
			var msg model.ChatMessage
			if err := ws.ReadJSON(&msg);err != nil{
				ws.WriteJSON(model.NoticeMessage{Message:"json error", Code:1})
			}else{
				broadcast <- msg
				ws.WriteJSON(model.NoticeMessage{Message:"success", Code:0})
			}
		}
	}else{
		ws.WriteJSON(model.NoticeMessage{Message:"token error", Code:1})
	}
}

func MessagePushHandle(){
	for{
		msg := <- broadcast
		fromUser := resource.GetUserByUsername(msg.From)
		toUser := resource.GetUserByUsername(msg.To)
		if fromUser.ID == 0 || toUser.ID == 0{
			log.Printf("消息发送失败")
			continue
		}
		log.Printf("收到来自%s发给%s的消息, 内容为:%s", fromUser.Name, toUser.Name, msg.Message)
		toClient := clients[toUser.Username]
		if toClient == nil{
			fmt.Println("存入redis")
			continue
		}
		err := toClient.WriteJSON(&msg)
		if err != nil{
			fmt.Println("push message err", err.Error())
		}
	}
}
