package models

import "time"

type GormModel struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"create_at,omitempty"`
	UpdatedAt time.Time `json:"update_at,omitempty"`
}