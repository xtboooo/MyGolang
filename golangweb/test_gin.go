/* package main

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context)  {
	// c.String(200, "hello, %s", "xtbo97")
	c.JSON(200, gin.H{
		"name": "xtbo97",
		"age": "20",
	})
}

func main() {
	e := gin.Default()
	e.GET("/hello", Hello)
	e.Run()
}
 */