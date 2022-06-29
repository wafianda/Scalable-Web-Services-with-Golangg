package controllers

import (
	database "Mygram/databases"
	"Mygram/helpers"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func PostPhotoByID(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Photo := models.Photo{}
	UserId := uint(userData["id"].(float64))

	ct := helpers.GetContentType(c)

	if ct == "application/json" {
		_ = c.ShouldBindJSON(&Photo)
	} else {
		_ = c.ShouldBind(&Photo)
	}

	Photo.UserId = UserId
	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{

			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": gin.H{
			"id":         Photo.ID,
			"title":      Photo.Title,
			"caption":    Photo.Caption,
			"photo_url":  Photo.PhotoUrl,
			"user_id":    Photo.UserID,
			"created_at": Photo.CreatedAt,
		}})
}

func ReadAllPhotoById(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	UserID := uint(userData["id"].(float64))

	var Photo []models.Photo
	var User models.User

	User.ID = UserID
	Email := fmt.Sprintf("%v", userData["email"])

	err := db.Debug().Where("user_id = ?", UserID).Find(&Photo).Error
	db.Select("username").Find(&User)

	var Photos []models.PhotoIncludeUserData
	var TempPhoto models.PhotoIncludeUserData

	for i := range Photo {
		TempPhoto.Id = Photo[i].ID
		TempPhoto.Title = Photo[i].Title
		TempPhoto.Caption = Photo[i].Caption
		TempPhoto.PhotoUrl = Photo[i].PhotoUrl
		TempPhoto.UserID = Photo[i].UserID
		TempPhoto.Created_at = Photo[i].CreatedAt
		TempPhoto.Updated_at = Photo[i].UpdatedAt
		Photos = append(Photos, TempPhoto)

	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   Photos,
	})
}

func UpdatePhotoByID(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	UserID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Photo := models.Photo{}

	idTemp, err := strconv.ParseUint(c.Param("photoId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Cant find idPhoto(1)",
		})
		return
	}

	Photo.ID = uint(idTemp)

	err = db.First(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Cant find idPhoto(2)",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	result := db.Model(&Photo).Where("user_id = ?", UserID).Updates(&Photo)

	log.Println(Photo)
	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "No row affected",
		})
		return
	}
	response := models.ResponseCommentUpdate{
		ID:        Photo.ID,
		UpdatedAt: Photo.UpdatedAt,
		Title:     Photo.Title,
		Caption:   Photo.Caption,
		PhotoUrl:  Photo.PhotoUrl,
		UserID:    Photo.UserID,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   response,
	})
}

func DeletePhotoById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	UserID := uint(userData["id"].(float64))

	Photo := models.Photo{}
	id := c.Param("photoId")

	result := db.Debug().Where("user_id = ?", UserID).Delete(&Photo, id)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Cant find idPhoto",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"message": "Your photo has been successfully deleted",
		},
	})
}
