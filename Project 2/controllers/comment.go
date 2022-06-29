package controllers

import (
	database "Mygram/databases"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GetCommentByID(c *gin.Context) {

	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   comments})
}

func PostCommentById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	requestComment := models.RequestComment{}
	userId := uint(userData["id"].(float64))
	c.ShouldBindJSON(&requestComment)

	_, err := govalidator.ValidateStruct(requestComment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	comment := models.Comment{
		UserID:    userId,
		Message:   requestComment.Message,
		PhotoID:   requestComment.PhotoID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = db.Create(&comment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": gin.H{

			"id":         comment.ID,
			"message":    comment.Message,
			"photo_id":   comment.PhotoID,
			"user_id":    comment.UserID,
			"created_at": comment.CreatedAt,
		}})
}

func UpdateCommentById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	requestComment := models.RequestComment{}
	commentId, _ := strconv.ParseUint(c.Param("commentId"), 10, 64)
	c.ShouldBindJSON(&requestComment)

	_, err := govalidator.ValidateStruct(requestComment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	comment := models.Comment{}
	err = db.First(&comment, "user_id = ? AND id = ?", userId, commentId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}

	comment = models.Comment{
		ID:        uint(commentId),
		Message:   requestComment.Message,
		UpdatedAt: time.Now(),
	}

	err = db.Model(&comment).Select("message", "updated_at").Updates(&comment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	photo := models.Photo{}
	err = db.First(&photo, "id = ?", userId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": gin.H{
			"id":         comment.ID,
			"title":      photo.Title,
			"caption":    photo.Caption,
			"photo_url":  photo.PhotoUrl,
			"user_id":    userId,
			"updated_at": comment.UpdatedAt,
		},
	})
}

func DeleteCommentById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	commentId, _ := strconv.ParseUint(c.Param("commentId"), 10, 64)
	comment := models.Comment{}
	err := db.First(&comment, "user_id = ? AND id = ?", userId, commentId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}
	comment = models.Comment{ID: uint(commentId)}
	err = db.Delete(&comment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"message": "Your comment has been successfully deleted"}})
}
