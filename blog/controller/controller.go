package controller

import (
	"blog/dao"
	"blog/models"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := models.User{
		Username: username,
		Password: password,
	}
	dao.Mgr.Register(&user)
	c.Redirect(301, "/")
}
