package route

import (
	"net/http"

	"example.com/goods-manage/common"
	"example.com/goods-manage/repository"
	"example.com/goods-manage/service"
	"github.com/gin-gonic/gin"
)

func getAllProducts(ctx *gin.Context) {
	repo := repository.NewProductRepo()
	service := service.NewProductService(repo)
	products, err := service.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.BaseResponse[any]{
			Code: http.StatusInternalServerError,
			Message: common.StatusGetFailed,
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, products)
}