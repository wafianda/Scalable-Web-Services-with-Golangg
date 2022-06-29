package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
)

var secret = "abcd"

func GenerateToken(id uint, email string) (signedToken string) {
	claims := jwt.MapClaims{
		"id": id,
		"email": email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ = token.SignedString([]byte(secret))

	return
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errCode := errors.New("please use correct authentication")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	if !bearer {
		return nil, errCode
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errCode
		}
		return []byte(secret), nil
	})

	if !token.Valid {
		return nil, errCode
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return nil, errCode
	}

	return token.Claims.(jwt.MapClaims), nil
}
