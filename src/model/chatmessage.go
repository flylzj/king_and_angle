package model

type ChatMessage struct {
	From		string		`json:"from"`
	To			string		`json:"to"`
	Message   	string		`json:"message"`
	Token    	string		`json:"token"`
	Type 		string		`json:"type"`
	Username	string		`json:"username"`
}

type NoticeMessage struct {
	Message  	string		`json:"message"`
	Code        uint		`json:"code"`
}

type PingMessage struct {
	Username	string		`json:"message"`
	Online		uint		`json:"online"`
}



