package chat

import (
	"config"
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
		config.Error.Println("connect err", err.Error())
		ctx.Abort()
	}

	defer ws.Close()

	var msg model.ChatMessage

	if err := ws.ReadJSON(&msg);err != nil{
		ws.WriteJSON(model.NoticeMessage{Message:"data error", Code:1})
	}
    config.Info.Println(msg)
	claims, result := resource.CheckToken(msg.Token)
	if result{
		currentUser := resource.GetUserById(claims.ID)
		clients[currentUser.Username] = ws

		for{
			var msg model.ChatMessage
			if err := ws.ReadJSON(&msg);err != nil{
				config.Error.Println("json error: ", err.Error())
				break
			}else{
				broadcast <- msg
				ws.WriteJSON(model.NoticeMessage{Message:"success", Code:0})
			}
		}
	}else{
		config.Error.Println("token error")
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
		config.Info.Printf("收到来自%s发给%s的消息, 内容为:%s", fromUser.Name, toUser.Name, msg.Message)
		toClient := clients[toUser.Username]
		if toClient == nil{
			config.Info.Printf("来自%s的消息, %s不在线, 消息被丢弃", fromUser.Name, toUser.Name)
			continue
		}
		if err := toClient.WriteJSON(&msg);err != nil{
			config.Error.Println("push message err", err.Error())
		}
		config.Info.Printf("%s发给%s的消息, 内容为:%s，转发成功", fromUser.Name, toUser.Name, msg.Message)
	}
}
