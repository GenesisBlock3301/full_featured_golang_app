package route

import (
	"bookshop/controllers"
	"bookshop/middlewares"

	"github.com/gin-gonic/gin"
)

func PostRoute(postRouter *gin.RouterGroup) {
	postRouter.GET("/",middlewares.JWTAuthMiddleware(), controller.AllPosts)
	postRouter.GET("/:postId",middlewares.JWTAuthMiddleware(), controller.FindByPostId)
	postRouter.POST("/",middlewares.JWTAuthMiddleware(), controller.InsertPost)
	postRouter.PUT("/:postId",middlewares.JWTAuthMiddleware(), controller.UpdatePost)
	postRouter.DELETE("/:postId",middlewares.JWTAuthMiddleware(), controller.DeleteByPostId)

}
