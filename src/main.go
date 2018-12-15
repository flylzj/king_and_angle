package main

import (
	"chat"
	"github.com/gin-gonic/gin"
	"resource"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()
	resource.User(route.Group("/api/user"))
	resource.KingAngAngle(route.Group("/api/king-and-angle", resource.JWTAuth()))
	resource.Blessing(route.Group("/api/blessing", resource.JWTAuth()))
	route.GET("/ws", chat.WsConnectionHandle)
	go chat.MessagePushHandle()
	route.Run("127.0.0.1:8003")
}