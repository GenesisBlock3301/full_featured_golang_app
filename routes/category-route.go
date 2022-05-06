package route

import (
	controller "bookshop/controllers"
	"bookshop/middlewares"

	"github.com/gin-gonic/gin"
)

func CategoruRoute(categoryRouter *gin.RouterGroup) {
	categoryRouter.GET("/", middlewares.JWTAuthMiddleware(), controller.AllCategories)
	categoryRouter.POST("/", middlewares.JWTAuthMiddleware(), controller.InsertCategory)
	categoryRouter.GET("/:cateryId", middlewares.JWTAuthMiddleware(), controller.CategoryById)
	categoryRouter.PUT("/:cateryId", middlewares.JWTAuthMiddleware(), controller.CategoryUpdate)
	categoryRouter.DELETE("/:cateryId", middlewares.JWTAuthMiddleware(), controller.CategoryDeleteById)
}
