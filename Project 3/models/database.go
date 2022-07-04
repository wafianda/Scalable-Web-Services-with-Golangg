package models

import (
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	var database *gorm.DB
	var err error
	if dsn == "" {
		database, err = gorm.Open(sqlite.Open("gorm.sqlite"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else {
		database, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
	}

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Board{}, &Column{}, &Ticket{})

	DB = database
}
