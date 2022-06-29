package models

type Comment struct {
	GormModel
	UserId  uint
	User    *User
	Photo   *Photo
	Content string `gorm:"not null" json:"content" form:"content" valid:"required~content is required"`
}

type RequestComment struct {
	Message string `json:"message" valid:"required"`
	PhotoID uint   `json:"photo_id,omitempty"`
}
