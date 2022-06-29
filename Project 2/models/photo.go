package models

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~title is required"`
	Caption  string `json:"description" form:"description"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~photo_url is required"`
	UserId   uint
	User     *User
}
