package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	router.GET("/register", controller.GoRegister)
	router.POST("/register", controller.Register)
	router.GET("/", controller.Index)
	router.Run(":8090")
}
