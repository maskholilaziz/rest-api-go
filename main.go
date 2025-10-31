package main

import (
	"jadilebihbaik/rest-api/controllers"
	"jadilebihbaik/rest-api/models"

	"github.com/gin-gonic/gin"
)

func main()  {
	// inisilisasi Gin
	router := gin.Default()

	// panggil koneksi database
	models.ConnectDatabase()

	// membuat router dengan method GET
	router.GET("/", func(c *gin.Context) {
		// return response JSON
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	// membuat route get all posts
	router.GET("/api/posts", controllers.FindPosts)

	// mulai server dengan port 3000
	router.Run(":3000")
}