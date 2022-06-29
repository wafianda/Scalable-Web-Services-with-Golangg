package models

type Product struct {
	GormModel
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	UserId      uint
	User        *User
}
