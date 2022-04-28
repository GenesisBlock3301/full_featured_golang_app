package route

import (
	"github.com/gin-gonic/gin"
)

func RootRoute(router *gin.Engine){
	// Static route
	router.Static("/media",".media")

	apiRouter := router.Group("api/v1")
	postRouter := apiRouter.Group("/posts")
	// Post route 
	PostRoute(postRouter)

	// Category route
	categoryRoute := apiRouter.Group("/categories")
	CategoruRoute(categoryRoute)

	// Auth router
	authRoute := apiRouter.Group("/auth")
	AuthRoute(authRoute)
	
}