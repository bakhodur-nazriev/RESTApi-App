package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restApi/models"
)

func GetUsers(ctx *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}
