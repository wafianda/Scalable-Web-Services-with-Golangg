package models

type Socialmedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~name is required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~social_media_url is required"`
	UserId         uint
	User           *User
}

type RequestSocialMedia struct {
	Name           string `json:"name" form:"name" valid:"required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_medial_url" valid:"required"`
}
