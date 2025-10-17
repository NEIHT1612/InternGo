package main

import (
	"example.com/goods-manage/db"
	"example.com/goods-manage/logger"
	"example.com/goods-manage/route"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	db.NewDB()

	logger.InitLogger()
	defer logger.Sync()

	server := gin.Default()
	server.Use(gin.Recovery())
	server.Use(logger.GinZapLogger(logger.Log))
	
	route.RegisterRoutes(server)

	server.Run(":8080")
}
