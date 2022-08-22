/* package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GoStatic(c *gin.Context) {
	c.HTML(http.StatusOK, "test_static.html", nil)
}

func main() {
	e := gin.Default()
	e.Static("assets", "./assets")
	e.LoadHTMLGlob("templates/*")
	e.GET("/goStatic", GoStatic)
	e.Run(":8081")
}
 */