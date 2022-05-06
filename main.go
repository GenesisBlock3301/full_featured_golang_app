package main

import (
	"bookshop/config"
	"bookshop/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// defer config.CloseDbConnection()
	config.ConnectWithDB()
	router := gin.Default()
	route.RootRoute(router)
	router.Run("localhost:8080")
}
