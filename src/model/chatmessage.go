package model

type ChatMessage struct {
	From		string		`json:"from"`
	To			string		`json:"to"`
	Message   	string		`json:"message"`
	SendTime	int			`json:"send_time"`
	Token    	string		`json:"token"`
	Type 		string		`json:"type"`
	Username	string		`json:"username"`
}

type NoticeMessage struct {
	Message  	string		`json:"message"`
	Code        uint		`json:"code"`
}

type PingMessage struct {
	Username	string		`json:"username"`
	Online		uint		`json:"online"`
}

type ChatMsg struct {
	ID 		uint	`grom:"AUTO_INCREMENT"`
	FromUsername		string		`gorm:"type:varchar(10)"`
	ToUsername			string		`gorm:"type:varchar(10)"`
	Message   	string		`gorm:"type:text"`
	SendTime	int
}


func (cm *ChatMsg) TableName() string{
	return "chat_msg"
}


