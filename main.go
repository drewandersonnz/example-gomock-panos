package main

import (
	"github.com/drpaneas/drew/controllers"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/ping", controllers.PingController)
	router.Run(":8080")
}
