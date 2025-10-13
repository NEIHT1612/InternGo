package route

import (
	"net/http"

	"example.com/goods-manage/models"
	"example.com/goods-manage/utils"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var customer models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse data"})
		return
	}
	if err := models.CreateUser(&customer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't create user"})
		return
	}
	ctx.JSON(http.StatusCreated, "Sign up successfully")
}

func login(ctx *gin.Context) {
	var customer models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse data"})
		return
	}
	
	if err := models.LoginUser(&customer); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		return
	}

	token, err := utils.GenerateToken(customer.Username, customer.CustomerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
