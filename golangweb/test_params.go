/* package main

import "github.com/gin-gonic/gin"

func testGet(c *gin.Context) {
	s := c.Query("username")
	pwd := c.DefaultQuery("password", "123456")
	c.String(200, "username: %s, password: %s", s, pwd)
}

func testPathParams(c *gin.Context) {
	s := c.Param("name")
	s2 := c.Param("age")
	c.String(200, "name: %s, age: %s", s, s2)
}

func goSearch(c *gin.Context) {
	c.HTML(200, "query.html", nil)
}

func Search(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	key := c.PostForm("key")
	c.String(200, "page: %s, key:%s", page, key)
}

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	// e.GET("/testGet", testGet)
	// e.GET("/testPath/:name/:age", testPathParams)
	e.GET("/goSearch", goSearch)
	e.POST("/search", Search)
	e.Run(":8888")
}
 */