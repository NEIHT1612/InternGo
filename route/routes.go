package route

import (
	"example.com/goods-manage/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/categories", getAllCategories)
	server.GET("/categories/:category_id", getCategoryByID)
	server.POST("categories", createCategory)
	server.PUT("/categories/:category_id", updateCategoryByID)
	server.DELETE("/categories/:category_id", deleteCategoryByID)

	server.GET("/products", getAllProducts)

	authenticated := server.Group("/")
	authenticated.Use(middleware.AuthMiddleware)
	{
		authenticated.POST("/upload", uploadFile)
		authenticated.POST("/uploadmultiple", UploadMultipleFiles)
	}

	server.POST("/signup", signup)
	server.POST("/login", login)
}
