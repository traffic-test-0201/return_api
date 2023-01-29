package main

import (
	"github.com/dongdonjuhahaha/traffic/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/returns", controllers.GetReturn)
	router.Run(":9090")
}
