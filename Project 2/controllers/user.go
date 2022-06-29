package controllers

import (
	database "Mygram/databases"
	"Mygram/helpers"
	"Mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserRegistration(c *gin.Context) {
	db := database.GetDB()

	var User models.User

	ct := helpers.GetContentType(c)

	if ct == "application/json" {
		_ = c.ShouldBindJSON(&User)
	} else {
		_ = c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"full_name": User.FullName,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()

	var User models.User

	ct := helpers.GetContentType(c)

	if ct == "application/json" {
		_ = c.ShouldBindJSON(&User)
	} else {
		_ = c.ShouldBind(&User)
	}

	password := User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePassword([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func UserUpdate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	UserID := uint(userData["id"].(float64))

	db.Where("id = ?", UserID).First(&User)

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.ID = UserID

	_, errUpdate := govalidator.ValidateStruct(User)

	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, gin.H{

			"error":   "Update Not Valid",
			"message": errUpdate,
		})
		return
	}

	err := db.Debug().Model(&User).Updates(models.User{Email: User.Email, Username: User.Username}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{

			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"id":         User.ID,
			"age":        User.Age,
			"email":      User.Email,
			"password":   User.Password,
			"username":   User.Username,
			"updated_at": User.UpdatedAt,
		},
	})
}

func UserDelete(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	UserID := uint(userData["id"].(float64))

	err := db.Where("id = ?", UserID).Delete(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   gin.H{"message": "Your account has been successfully deleted"},
	})
}
