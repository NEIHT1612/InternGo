package route

import (
	"net/http"

	"example.com/goods-manage/common"
	"example.com/goods-manage/models"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var customer models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BaseResponse[any]{
			Code: http.StatusBadRequest,
			Message: common.StatusBadRequest,
			Data: nil,
		})
		return
	}
	if err := models.CreateUser(&customer); err != nil {
		ctx.JSON(http.StatusInternalServerError, common.BaseResponse[any]{
			Code: http.StatusInternalServerError,
			Message: common.StatusCreateFailed,
			Data: nil,
		})
		return
	}
	
	ctx.JSON(http.StatusOK, common.BaseResponse[any]{
		Code: http.StatusOK,
		Message: "Signup successfully",
		Data: nil,
	})
}

func login(ctx *gin.Context) {
	var customer models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BaseResponse[any]{
			Code: http.StatusBadRequest,
			Message: common.StatusBadRequest,
			Data: nil,
		})
		return
	}

	if err := models.LoginUser(&customer); err != nil {
		ctx.JSON(http.StatusUnauthorized, common.BaseResponse[any]{
			Code: http.StatusUnauthorized,
			Message: common.StatusUnauthorized,
			Data: nil,
		})
		return
	}

	token, err := common.GenerateToken(customer.Username, customer.CustomerID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.BaseResponse[any]{
			Code: http.StatusInternalServerError,
			Message: common.StatusCreateFailed,
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, common.BaseResponse[any]{
		Code: http.StatusOK,
		Message: common.StatusCreateSuccess,
		Data: map[string]any{"token": token},
	})
}
