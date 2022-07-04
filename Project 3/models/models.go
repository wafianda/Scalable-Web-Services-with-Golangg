package models

import "time"

type Board struct {
	BaseModel
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	StartDate time.Time `json:"start_date" gorm:"not null"`
	EndDate   time.Time `json:"end_date" gorm:"not null"`
	Tickets   []Ticket  `json:"tickets" gorm:"foreignkey:BoardID"`
}

type Column struct {
	BaseModel
	Name    string   `json:"name" gorm:"type:varchar(255);not null"`
	Order   uint     `json:"order" gorm:"type:integer;not null;default:0;unique;"`
	Tickets []Ticket `json:"tickets" gorm:"foreignkey:ColumnID"`
}

type Ticket struct {
	BaseModel
	Title       string `json:"title" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:varchar(1025)"`
	BoardID     uint   `json:"board_id" gorm:"type:integer"`
	Board       Board  `json:"board" gorm:"foreignkey:BoardID"`
	ColumnID    uint   `json:"column_id" gorm:"type:integer; not null"`
	Column      Column `json:"column" gorm:"foreignkey:ColumnID"`
}
