package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func testMW(c *gin.Context) {
	c.String(200, "hello %s", "xtbo")
}

func myMiddleWare1(c *gin.Context) {
	fmt.Println("我的第一个中间件")
}

func myMiddleWare2(c *gin.Context) {
	fmt.Println("我的第二个中间件")
}

func main() {
	e := gin.Default()
	e.Use(myMiddleWare1, myMiddleWare2)
	e.GET("/testmw", testMW)
	e.Run(":8084")
}
