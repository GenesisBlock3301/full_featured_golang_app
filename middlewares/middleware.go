package middlewares

import (
	"bookshop/services"
	"net/http"
	"github.com/gin-gonic/gin"
)



func JWTAuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		err := services.TokenValid(ctx)
		if err != nil{
			ctx.String(http.StatusUnauthorized,"User not authorized!")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
	
}