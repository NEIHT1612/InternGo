package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func NotFound(c *gin.Context){
	c.JSON(http.StatusNotFound, BaseResponse[any]{
		Code: http.StatusNotFound,
		Message: "Route not found",
		Data: nil,
	})
}