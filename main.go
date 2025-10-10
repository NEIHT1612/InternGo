package main

import (
	"example.com/goods-manage/db"
	"example.com/goods-manage/route"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	db.NewDB()

	server := gin.Default()

	route.RegisterRoutes(server)

	server.Run(":8080")
}
