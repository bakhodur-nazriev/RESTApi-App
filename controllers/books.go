package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restApi/models"
)

func GetBooks(ctx *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{"data": books})
}
