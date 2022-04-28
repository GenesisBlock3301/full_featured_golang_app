package route

import (
	controller "bookshop/controllers"
	"bookshop/middlewares"

	"github.com/gin-gonic/gin"
)



func AuthRoute(authRouter *gin.RouterGroup){
	authRouter.POST("/register",controller.Register)
	authRouter.POST("/login",controller.Login)
	authRouter.GET("/user",middlewares.JWTAuthMiddleware(),controller.CurrentUser)
}