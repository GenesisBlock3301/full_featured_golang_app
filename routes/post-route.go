package route

import (
	"bookshop/controllers"

	"github.com/gin-gonic/gin"
)

func PostRoute(postRouter *gin.RouterGroup){
	postRouter.GET("/",controller.AllPosts)
	postRouter.GET("/:postId",controller.FindByPostId)
	postRouter.POST("/",controller.InsertPost)
	postRouter.PUT("/:postId",controller.UpdatePost)
	postRouter.DELETE("/:postId",controller.DeleteByPostId)

}
