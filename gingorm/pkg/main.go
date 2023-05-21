package main

import (
	"log"

	"github.com/Shrut26/gingorm/config"
	"github.com/Shrut26/gingorm/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	log.Fatal(router.Run(":8080"))
}
