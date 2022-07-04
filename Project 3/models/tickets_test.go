package models_test

import (
	"kanban-board/models"
	"testing"
)

func TestTickets(t *testing.T) {
	models.DB = MockDb()
	var ticket models.Ticket
	// Test Create
	t.Run("Create ticket", func(t *testing.T) {
		ticket = models.Ticket{
			Title: "Test ticket",
		}
		if err := models.TicketManager.Create(&ticket); err != nil {
			t.Errorf("Error creating ticket: %v", err)
		}
	})
	// Test Update
	t.Run("Update ticket", func(t *testing.T) {
		ticket.Title = "Test ticket updated"
		if err := models.TicketManager.Update(&ticket, uint64(ticket.BaseModel.ID)); err != nil {
			t.Errorf("Error updating ticket: %v", err)
		}
	})
	// Test GetById
	t.Run("Get ticket by id", func(t *testing.T) {
		ticket, err := models.TicketManager.GetById(1)
		if err != nil {
			t.Errorf("Error getting ticket by id: %v", err)
		}
		if ticket.Title != "Test ticket updated" {
			t.Errorf("Error getting ticket by id: %d and title %s, got %s", 1, "Test ticket updated", ticket.Title)
		}
	})
	// Test GetAll
	t.Run("Get all tickets", func(t *testing.T) {
		tickets, err := models.TicketManager.GetAll()
		if err != nil {
			t.Errorf("Error getting all tickets: %v", err)
		}
		if len(*tickets) != 1 {
			t.Errorf("Error getting all tickets: expected 1 ticket, got %d", len(*tickets))
		}
		if (*tickets)[0].Title != "Test ticket updated" {
			t.Errorf("Error getting all tickets: expected ticket title Test ticket updated, got %s", (*tickets)[0].Title)
		}
	})
	// Test DeleteById
	t.Run("Delete ticket by id", func(t *testing.T) {
		if err := models.TicketManager.DeleteById(1); err != nil {
			t.Errorf("Error deleting ticket by id: %v", err)
		}
	})
}
