package middleware

import (
	"example.com/goods-manage/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(401, gin.H{"error": "Authorization header is required"})
		context.Abort()
		return
	}

	customerID, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(401, gin.H{"error": "Invalid token"})
		context.Abort()
		return
	}
	
	context.Set("customer_id", customerID)
	context.Next()
}