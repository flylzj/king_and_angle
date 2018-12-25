package model

import (
	"config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db *gorm.DB
)


func init(){
	var err error
	Db, err = gorm.Open("mysql", config.MYSQLURI)
	if err != nil{
		config.Error.Fatalln("数据库链接错误")
	}

	if !Db.HasTable(&User{}){
		Db.CreateTable(&User{})
	}

	if !Db.HasTable(&ChatMsg{}){
		Db.CreateTable(&ChatMsg{})
	}

	Db.AutoMigrate(&User{}, &ChatMsg{})
	return
}