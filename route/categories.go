package route

import (
	"net/http"

	"example.com/goods-manage/common"
	"example.com/goods-manage/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getAllCategories(ctx *gin.Context) {
	categories, err := models.GetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.BaseResponse[any]{
			Code: http.StatusInternalServerError,
			Message: common.StatusGetFailed,
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func getCategoryByID(ctx *gin.Context) {
	categoryID := uuid.MustParse(ctx.Param("category_id"))
	category, err := models.GetCategoryByID(categoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.BaseResponse[any]{
			Code: http.StatusInternalServerError,
			Message: common.StatusGetFailed,
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func createCategory(ctx *gin.Context) {
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BaseResponse[any]{
			Code: http.StatusBadRequest,
			Message: common.StatusBadRequest,
			Data: nil,
		})
		return
	}
	
	if err := models.CreateCategory(&category); err != nil {
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
		Data: map[string]any{"category": category},
	})
}

func updateCategoryByID(ctx *gin.Context) {
	categoryID := uuid.MustParse(ctx.Param("category_id"))
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BaseResponse[any]{
			Code: http.StatusBadRequest,
			Message: common.StatusBadRequest,
			Data: nil,
		})
		return
	}
	category.CategoryID = categoryID
	if err := models.UpdateCategoryByID(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, common.BaseResponse[any]{
			Code: http.StatusInternalServerError,
			Message: common.StatusUpdateFailed,
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, common.BaseResponse[any]{
		Code: http.StatusOK,
		Message: common.StatusUpdateSuccess,
		Data: map[string]any{"category": category},
	})
}

func deleteCategoryByID(ctx *gin.Context) {
	categoryID := uuid.MustParse(ctx.Param("category_id"))

	if err := models.DeleteCategoryByID(categoryID); err != nil {
		ctx.JSON(http.StatusInternalServerError, common.BaseResponse[any]{
			Code: http.StatusInternalServerError,
			Message: common.StatusDeleteFailed,
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, common.BaseResponse[any]{
		Code: http.StatusOK,
		Message: common.StatusDeleteSuccess,
		Data: nil,
	})
}
