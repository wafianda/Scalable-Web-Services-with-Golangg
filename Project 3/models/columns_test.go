package models_test

import (
	"kanban-board/models"
	"testing"
)

func TestColumns(t *testing.T) {
	models.DB = MockDb()
	var column models.Column
	// Test Create
	t.Run("Create column", func(t *testing.T) {
		column = models.Column{
			Name: "Test column",
		}
		if err := models.ColumnManager.Create(&column); err != nil {
			t.Errorf("Error creating column: %v", err)
		}
	})
	// Test Update
	t.Run("Update column", func(t *testing.T) {
		column.Name = "Test column updated"
		if err := models.ColumnManager.Update(&column, uint64(column.BaseModel.ID)); err != nil {
			t.Errorf("Error updating column: %v", err)
		}
	})
	// Test GetById
	t.Run("Get column by id", func(t *testing.T) {
		column, err := models.ColumnManager.GetById(1)
		if err != nil {
			t.Errorf("Error getting column by id: %v", err)
		}
		if column.Name != "Test column updated" {
			t.Errorf("Error getting column by id: %d and name %s, got %s", 1, "Test column updated", column.Name)
		}
	})
	// Test GetAll
	t.Run("Get all columns", func(t *testing.T) {
		columns, err := models.ColumnManager.GetAll()
		if err != nil {
			t.Errorf("Error getting all columns: %v", err)
		}
		if len(*columns) != 1 {
			t.Errorf("Error getting all columns: expected 1 column, got %d", len(*columns))
		}
		if (*columns)[0].Name != "Test column updated" {
			t.Errorf("Error getting all columns: expected column name Test column updated, got %s", (*columns)[0].Name)
		}
	})
	// Test DeleteById
	t.Run("Delete column by id", func(t *testing.T) {
		if err := models.ColumnManager.DeleteById(1); err != nil {
			t.Errorf("Error deleting column by id: %v", err)
		}
	})
}
