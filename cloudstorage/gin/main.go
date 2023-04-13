package main

import (
	"cloudstorage/gin/handler"
	
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("pages/*")
	r.Static("/static", "./static")

	r.GET("/index", handler.IndexHandler)
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	r.POST("/login", handler.LoginHandler)

	r.GET("/regist", func(c *gin.Context) {
		c.HTML(200, "regist.html", nil)
	})
	r.POST("/regist", handler.RegistHandler)
	r.POST("/verify", handler.VerifyHandler)
	r.GET("/logout", handler.LogoutHandler)

	r.POST("/upload", handler.UploadHandler)
	r.GET("/download", handler.DownloadHandler)

	r.GET("/registVIP", func(c *gin.Context) {
		c.HTML(200, "registVIP.html", nil)
	})
	r.POST("/registVIP", handler.RegistVIPHandler)

	r.GET("/delete", handler.DeleteFileHandler)

	r.Run(":8080")
}