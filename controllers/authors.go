package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restApi/models"
)

type CreateAuthorInput struct {
	Name        string `json:"name" binding:"name"`
	Bio         string `json:"bio" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
	DateOfDeath string `json:"date_of_death" binding:"required"`
	Image       string `json:"image" binding:"required"`
}

type UpdateAuthorInput struct {
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	DateOfBirth string `json:"date_of_birth"`
	DateOfDeath string `json:"date_of_death"`
	Image       string `json:"image"`
}

func GetAuthors(ctx *gin.Context) {
	var authors []models.Author
	models.DB.Find(&authors)

	ctx.JSON(http.StatusOK, gin.H{"data": authors})
}

func GetAuthor(ctx *gin.Context) {
	var author models.Author

	if err := models.DB.Where(":id = ?", ctx.Param("id")).First(&author).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": author})
}

func CreateAuthor(ctx *gin.Context) {
	/* Validate author */
	var input CreateAuthorInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/* Create author */
	author := models.Author{Name: input.Name, Bio: input.Bio, DateOfBirth: input.DateOfBirth, DateOfDeath: input.DateOfDeath, Image: input.Image}
	models.DB.Create(&author)
	ctx.JSON(http.StatusOK, gin.H{"data": author})
}

func UpdateAuthor(ctx *gin.Context) {
	/* Get model if exist */
	var author models.Author
	if err := models.DB.Where("id = ?", ctx.Param("id")).Find(&author).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/* Validate input */
	var input UpdateAuthorInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&author).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"data": author})
}

func DeleteAuthor(ctx *gin.Context) {
	/* Get model if exist */
	var author models.Author
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&author).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&author)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
