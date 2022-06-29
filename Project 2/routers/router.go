package routers

import (
	"Mygram/controllers"
	"Mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func InitApplication() *gin.Engine {
	r := gin.Default()
	// user/register
	// user/login
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", controllers.UserRegistration)
		userGroup.POST("/login", controllers.UserLogin)
		userGroup.PUT("/", middlewares.Authentication(), controllers.UserUpdate)
		userGroup.DELETE("/", middlewares.Authentication(), controllers.UserDelete)
	}

	commentGroup := r.Group("/comment")
	{
		commentGroup.POST("/", controllers.PostCommentByID)
		commentGroup.GET("/", controllers.GetCommentByID)
		commentGroup.PUT("/:commentId", controllers.UpdateCommentByID)
		commentGroup.DELETE("/:commentId", controllers.DeleteCommentByID)
	}

	photoGroup := r.Group("/photo")
	{
		photoGroup.POST("/", controllers.PostPhotoByID)
		photoGroup.GET("/", controllers.ReadPhotoByID)
		photoGroup.GET("/:photoId", controllers.ReadAllPhotoByID)
		photoGroup.PUT("/:photoId", controllers.UpdatePhotoByID)
		photoGroup.DELETE("/:photoId", controllers.DeletePhotoByID)
	}

	socialmediaGroup := r.Group("/socialmedia")
	{
		socialmediaGroup.POST("/", controllers.PostSocialmediaByID)
		socialmediaGroup.GET("/", controllers.GetSocialmediaByID)
		socialmediaGroup.PUT("/:socialmediaId", controllers.UpdateSocialmediaByID)
		socialmediaGroup.DELETE("/:socialmediaId", controllers.DeleteSocialmediaByID)
	}

	return r
}
