package main

import (
	"github.com/gin-gonic/gin"
	"github.com/dongdonjuhahaha/traffic/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/returns", controllers.GetReturn )
	router.Run("localhost:9090")
}
