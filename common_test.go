package main

import (
	"bookshop/middlewares"

	"github.com/gin-gonic/gin"
)

func getRouter(withTemplate bool) *gin.Engine{
	r := gin.Default()
	if withTemplate{
		r.Use(middlewares.JWTAuthMiddleware())
	}
	return r
}


