package serializers

import "kanban-board/models"

type TicketBody struct {
	Title       string `json:"title" binding:"required" validate:"required"`
	Description string `json:"description"`
	BoardID     uint   `json:"board_id" binding:"required" validate:"required"`
	ColumnID    uint   `json:"column_id" binding:"required" validate:"required"`
}

func (t *TicketBody) ToModel() models.Ticket {
	return models.Ticket{
		Title:       t.Title,
		Description: t.Description,
		BoardID:     t.BoardID,
		ColumnID:    t.ColumnID,
	}
}
