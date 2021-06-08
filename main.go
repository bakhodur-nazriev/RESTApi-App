package main

import (
	"github.com/gin-gonic/gin"
	"restApi/controllers"
	"restApi/models"
)

func main() {
	req := gin.Default()

	models.ConnectDatabase()

	/* Users Routes */
	req.GET("/users", controllers.GetUsers)
	req.GET("/users/:id", controllers.GetUser)
	req.POST("/users", controllers.CreateUser)
	req.POST("/users/:id", controllers.UpdateUser)
	req.DELETE("/users/:id", controllers.DeleteUser)

	/* Books Routes */
	req.GET("/books", controllers.GetBooks)
	req.GET("/books/:id", controllers.GetBook)
	req.POST("/books", controllers.CreateBook)
	req.PATCH("/books/:id", controllers.UpdateBook)
	req.DELETE("/books/:id", controllers.DeleteBook)

	/* Authors Routes */
	req.GET("/authors", controllers.GetAuthors)
	req.GET("/authors/:id", controllers.GetAuthor)
	req.POST("/authors", controllers.CreateAuthor)
	req.PATCH("/authors/:id", controllers.UpdateAuthor)
	req.DELETE("/authors/:id", controllers.DeleteAuthor)

	req.Run(":8080")
}
