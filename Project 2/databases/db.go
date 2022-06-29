package database

import (
	"Mygram/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "127.0.0.1"
	user     = "postgres"
	password = "postgres"
	port     = "5432"
	dbName   = "go_api_jwt"
	db       *gorm.DB
	err      error
)

func InitDB() {
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)
	db, err = gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		log.Panic("Database error", err.Error())
	}
	log.Println("Database connected")
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
