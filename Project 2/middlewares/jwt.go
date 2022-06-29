package middlewares

import (
	database "Mygram/databases"
	"Mygram/helpers"
	"Mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": err.Error(),
			})
		}

		c.Set("user_data", verifyToken)
		c.Next()
	}
}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		productId, err := strconv.Atoi(c.Param("productId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}

		userData := c.MustGet("user_data").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		Product := models.Product{}

		err = db.Debug().Select("user_id").First(&Product, uint(productId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if Product.UserId != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "don't have access",
			})
			return
		}

		c.Next()
	}
}
