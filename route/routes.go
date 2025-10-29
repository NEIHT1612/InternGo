package route

import (
	"example.com/goods-manage/common"
	"example.com/goods-manage/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	categories := server.Group(("/categories"))
	{
		categories.GET("", getAllCategories)
		categories.GET("/:category_id", getCategoryByID)
		categories.POST("", createCategory)
		categories.PUT("/:category_id", updateCategoryByID)
		categories.DELETE("/:category_id", deleteCategoryByID)
	}

	server.GET("/products", getAllProducts)

	authenticated := server.Group("/")
	authenticated.Use(middleware.AuthMiddleware)
	{
		authenticated.POST("/upload", uploadFile)
		authenticated.POST("/uploadmultiple", UploadMultipleFiles)
	}

	server.POST("/signup", signup)
	server.POST("/login", login)

	server.NoRoute(common.NotFound)
}
