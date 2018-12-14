package model

import (
	"config"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db *gorm.DB
)


func init(){
	var err error
	Db, err = gorm.Open("mysql", config.MYSQLURI)
	if err != nil{
		fmt.Println(err)
		fmt.Println("数据库链接错误")
		os.Exit(0)
	}

	if !Db.HasTable(&User{}){
		Db.CreateTable(&User{})
	}

	Db.AutoMigrate(&User{})
	return
}