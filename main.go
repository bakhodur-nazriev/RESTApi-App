package main

import (
	"github.com/gin-gonic/gin"
	"restApi/controllers"
	"restApi/models"
)

func main() {
	req := gin.Default()

	models.ConnectDatabase()

	req.GET("/", controllers.GetUsers)
	req.GET("/books", controllers.GetBooks)

	req.Run(":8080")
}
