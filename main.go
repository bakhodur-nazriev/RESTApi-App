package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restApi/models"
)

func main() {
	req := gin.Default()

	models.ConnectDatabase()

	req.GET("/", users)
	req.GET("/books", books)

	req.Run(":8080")
}

func users(ctx *gin.Context) {
	user := struct {
		id    int
		name  string
		age   int
		phone string
	}{1, "Bakhodur Nazriev", 25, "+992-985-30-52-55"}

	ctx.JSON(http.StatusOK, gin.H{
		"id":    user.id,
		"name":  user.name,
		"age":   user.age,
		"phone": user.phone,
	})
}

func books(ctx *gin.Context) {

}
