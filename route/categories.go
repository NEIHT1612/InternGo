package route

import (
	"net/http"

	"example.com/goods-manage/models"
	"github.com/gin-gonic/gin"
)

func getAllCategories(ctx *gin.Context) {
	categories, err := models.GetAllCategories()
	println(categories)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't get data category"})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}
