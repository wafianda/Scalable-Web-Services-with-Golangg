package main

import (
	"kanban-board/controllers"
	docs "kanban-board/docs"
	"kanban-board/models"

	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	godotenv.Load()

	models.ConnectDatabase()
	r := gin.Default()
	// r.Use(CorsMiddleware())
	docs.SwaggerInfo.BasePath = "/"

	boards := r.Group("/boards")
	{
		boards.GET("/", controllers.GetBoards())
		boards.POST("/", controllers.CreateBoard())
		boards.GET("/:id", controllers.GetBoard())
		boards.PUT("/:id", controllers.UpdateBoard())
		boards.DELETE("/:id", controllers.DeleteBoard())
	}
	columns := r.Group("/columns")
	{
		columns.GET("/", controllers.GetColumns())
		columns.POST("/", controllers.CreateColumn())
		columns.GET("/:id", controllers.GetColumn())
		columns.PUT("/:id", controllers.UpdateColumn())
		columns.DELETE("/:id", controllers.DeleteColumn())
	}
	tickets := r.Group("/tickets")
	{
		tickets.GET("/", controllers.GetTickets())
		tickets.POST("/", controllers.CreateTicket())
		tickets.GET("/:id", controllers.GetTicket())
		tickets.PUT("/:id", controllers.UpdateTicket())
		tickets.DELETE("/:id", controllers.DeleteTicket())
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
