package serializers

import "kanban-board/models"

type ColumnBody struct {
	Name  string `json:"name" binding:"required" validate:"required,min=3,max=255"`
	Order uint   `json:"order" binding:"required" validate:"required"`
}

func (body *ColumnBody) ToModel() models.Column {
	return models.Column{
		Name:  body.Name,
		Order: body.Order,
	}
}
