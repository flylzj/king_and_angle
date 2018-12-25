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
		return
	}

	defer ws.Close()

	var msg model.ChatMessage

	if err := ws.ReadJSON(&msg);err != nil{
		ws.WriteJSON(model.NoticeMessage{Message:"data error", Code:1})
		return
	}
	claims, result := resource.CheckToken(msg.Token)
	if result{
		currentUser := resource.GetUserById(claims.ID)
		clients[currentUser.Username] = ws
		messages := resource.GetMessage(currentUser.Username)
		for _, message := range messages{
			config.Info.Println(message)
			if err := ws.WriteJSON(&message); err != nil{
				config.Error.Println("消息发送失败", err.Error())
			}
		}
		for{
			var msg model.ChatMessage
			if err := ws.ReadJSON(&msg);err != nil{
				config.Error.Println("json error: ", err.Error())
				err := ws.WriteJSON(model.NoticeMessage{Message:"json error", Code:1})
				if err != nil{
					config.Error.Println("send message error:", err.Error())
					delete(clients, currentUser.Username)
					break
				}
				continue
			}
			if msg.Type == "ping"{
				if clients[msg.Username] != nil{
					err := ws.WriteJSON(model.PingMessage{Username:msg.Username, Online:0})
					if err != nil{
						config.Error.Println("send message error:", err.Error())
						break
					}
				}else{
					err := ws.WriteJSON(model.PingMessage{Username:msg.Username, Online:1})
					if err != nil {
						config.Error.Println("send message error:", err.Error())
						break
					}
				}
			}else{
				broadcast <- msg
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
		if resource.SaveMessage(msg){
			config.Info.Println("消息保存成功")
		}else{
			config.Info.Println("消息保存失败")
		}
		config.Info.Printf("收到来自%s发给%s的消息, 内容为:%s", fromUser.Name, toUser.Name, msg.Message)
		toClient := clients[toUser.Username]
		fromClient := clients[fromUser.Username]
		if toClient == nil{
			//err := fromClient.WriteJSON(model.NoticeMessage{Message:"消息发送失败,对方不在线", Code:1})
			//if err != nil{
			//	config.Error.Println("send message error : ", err.Error())
			//}
			//config.Info.Printf("来自%s的消息, %s不在线, 消息被丢弃", fromUser.Name, toUser.Name)
			continue
		}else {
			err := fromClient.WriteJSON(model.NoticeMessage{Message:"消息发送成功", Code:0})
			if err != nil{
				config.Error.Println("send message error : ", err.Error())
			}
		}
		if err := toClient.WriteJSON(&msg);err != nil{
			config.Error.Println("push message err", err.Error())
			continue
		}
		config.Info.Printf("%s发给%s的消息, 内容为:%s，转发成功", fromUser.Name, toUser.Name, msg.Message)
	}
}
