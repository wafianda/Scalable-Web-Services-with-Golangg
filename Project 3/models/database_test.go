package models_test

import (
	"kanban-board/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MockDb() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&models.Board{}, &models.Column{}, &models.Ticket{})
	return database
}
