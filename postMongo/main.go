package main

import (
	"github.com/Shrut26/Postmongo/config"
	"github.com/Shrut26/Postmongo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoutes(router)
	router.Run(":8080")
}
