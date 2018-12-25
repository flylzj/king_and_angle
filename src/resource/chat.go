package resource

import (
	"config"
	"model"
)

func SaveMessage(message model.ChatMessage) bool{
	if err := model.Db.Create(&model.ChatMsg{FromUsername:message.From, ToUsername:message.To, Message:message.Message, SendTime:message.SendTime}).Error; err != nil{
		config.Error.Println("save message error", err.Error())
		return false
	}
	return true
}

func GetMessage(username string) []model.ChatMessage{
	var messages []model.ChatMsg
	var msgs []model.ChatMessage
	if err := model.Db.Where("from_username=? or to_username=?", username, username).Find(&messages).Error; err!= nil{
		config.Error.Println("get message error", err.Error())
		return msgs
	}
	for _, message := range messages{
		msgs = append(msgs, model.ChatMessage{From:message.FromUsername, To:message.ToUsername, Message:message.Message, SendTime:message.SendTime})
	}
	return msgs
}
