package config

import (
	"io"
	"log"
	"os"
)

var(
	MYSQLURI string
	Info *log.Logger
	Error *log.Logger
)

func init(){
	errFile, err := os.OpenFile("err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil{
		log.Fatalln("打开文件失败", err.Error())
	}
	MYSQLURI = "root:123456@(localhost)/christmas?charset=utf8&parseTime=True&loc=Local"
	Info = log.New(os.Stdout, "Info:", log.Ldate | log.Ltime | log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stdout, errFile), "Error:", log.Ldate | log.Ltime | log.Llongfile)
}
