package route

import (
	"net/http"

	"example.com/goods-manage/repository"
	"example.com/goods-manage/service"
	"github.com/gin-gonic/gin"
)

func getAllProducts(ctx *gin.Context) {
	repo := repository.NewProductRepo()
	service := service.NewProductService(repo)
	products, err := service.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't get data products"})
		return
	}
	ctx.JSON(http.StatusOK, products)
}