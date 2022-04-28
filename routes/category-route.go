package route

import (
	controller "bookshop/controllers"

	"github.com/gin-gonic/gin"
)


func CategoruRoute(categoryRouter *gin.RouterGroup){
	categoryRouter.GET("/",controller.AllCategories)
	categoryRouter.POST("/",controller.InsertCategory)
	categoryRouter.GET("/:cateryId",controller.CategoryById)
	categoryRouter.PUT("/:cateryId",controller.CategoryUpdate)
	categoryRouter.DELETE("/:cateryId",controller.CategoryDeleteById)
}