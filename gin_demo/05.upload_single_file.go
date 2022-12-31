package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func upload_file(r *gin.Engine) {
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}
		// c.JSON(200, gin.H{"message": file.Header.Context})
		c.SaveUploadedFile(file, "./data/"+file.Filename)
		c.String(http.StatusOK, file.Filename)
	})
}

func upload_file_2(r *gin.Engine) {
	r.POST("/upload2", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file: %v", err)
		}
		//headers.Size 获取文件大小
		if headers.Size > 1024*1024*2 {
			fmt.Println("文件太大了")
			return
		}
		//headers.Header.Get("Content-Type")获取上传文件的类型
		if headers.Header.Get("Content-Type") != "image/png" {
			fmt.Println("只允许上传png图片")
			return
		}
		c.SaveUploadedFile(headers, "./data/"+headers.Filename)
		c.String(http.StatusOK, headers.Filename)
	})
}

func main() {
	r := gin.Default()
	r.Run()
	upload_file(r)
	upload_file_2(r)
}
