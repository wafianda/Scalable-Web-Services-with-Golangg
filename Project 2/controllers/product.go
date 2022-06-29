package controllers

import (
	database "Mygram/databases"
	"Mygram/helpers"
	"Mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()

	var Product models.Product

	ct := helpers.GetContentType(c)

	if ct == "application/json" {
		_ = c.ShouldBindJSON(&Product)
	} else {
		_ = c.ShouldBind(&Product)
	}

	userData := c.MustGet("user_data").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	Product.UserId = userId

	err := db.Debug().Create(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Product)
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	ct := helpers.GetContentType(c)
	userData := c.MustGet("user_data").(jwt.MapClaims)

	userId := uint(userData["id"].(float64))
	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))

	if ct == "application/json" {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserId = userId
	Product.ID = uint(productId)

	err := db.Model(&Product).Where("id = ?", productId).Updates(models.Product{
		Title:       Product.Title,
		Description: Product.Description,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}
