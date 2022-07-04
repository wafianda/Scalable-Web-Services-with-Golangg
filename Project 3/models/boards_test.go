package models_test

import (
	"kanban-board/models"
	"testing"
	"time"
)

func TestBoards(t *testing.T) {
	models.DB = MockDb()
	var board models.Board
	// Test Create
	t.Run("Create board", func(t *testing.T) {
		board = models.Board{
			Name:      "Test board",
			StartDate: time.Now(),
			EndDate:   time.Now(),
		}
		if err := models.BoardManager.Create(&board); err != nil {
			t.Errorf("Error creating board: %v", err)
		}
	})

	// Test Update
	t.Run("Update board", func(t *testing.T) {
		board.Name = "Test board updated"
		if err := models.BoardManager.Update(&board, uint64(board.BaseModel.ID)); err != nil {
			t.Errorf("Error updating board: %v", err)
		}
	})

	// Test GetById
	t.Run("Get board by id", func(t *testing.T) {
		board, err := models.BoardManager.GetById(1)
		if err != nil {
			t.Errorf("Error getting board by id: %v", err)
		}
		if board.Name != "Test board updated" {
			t.Errorf("Error getting board by id: %d and name %s, got %s", 1, "Test board updated", board.Name)
		}
	})

	// Test GetAll
	t.Run("Get all boards", func(t *testing.T) {
		boards, err := models.BoardManager.GetAll()
		if err != nil {
			t.Errorf("Error getting all boards: %v", err)
		}
		if len(*boards) != 1 {
			t.Errorf("Error getting all boards: expected 1 board, got %d", len(*boards))
		}
		if (*boards)[0].Name != "Test board updated" {
			t.Errorf("Error getting all boards: expected board name Test board updated, got %s", (*boards)[0].Name)
		}
	})

	// Test DeleteById
	t.Run("Delete board by id", func(t *testing.T) {
		if err := models.BoardManager.DeleteById(1); err != nil {
			t.Errorf("Error deleting board by id: %v", err)
		}
		boards, err := models.BoardManager.GetAll()
		if err != nil {
			t.Errorf("Error getting all boards: %v", err)
		}
		if len(*boards) != 0 {
			t.Errorf("Error getting all boards: expected 0 boards, got %d", len(*boards))
		}
	})

}
