package route

import (
	"net/http"

	"example.com/goods-manage/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getAllCategories(ctx *gin.Context) {
	categories, err := models.GetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't get data category"})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func getCategoryByID(ctx *gin.Context) {
	categoryID := uuid.MustParse(ctx.Param("category_id"))
	category, err := models.GetCategoryByID(categoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't get data category"})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func createCategory(ctx *gin.Context) {
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse data"})
		return
	}
	
	if err := models.CreateCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't create category"})
		return
	}
	ctx.JSON(http.StatusCreated, category)
}

func updateCategoryByID(ctx *gin.Context) {
	categoryID := uuid.MustParse(ctx.Param("category_id"))
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse data"})
		return
	}
	category.CategoryID = categoryID
	if err := models.UpdateCategoryByID(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't update category"})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func deleteCategoryByID(ctx *gin.Context) {
	categoryID := uuid.MustParse(ctx.Param("category_id"))
	if err := models.DeleteCategoryByID(categoryID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't delete category"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
