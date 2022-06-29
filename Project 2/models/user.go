package models

import (
	"Mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string  `gorm:"not null" json:"user_name" form:"user_name" valid:"required~user_name is required"`
	Email    string  `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required,email~email invalid"`
	Password string  `gorm:"not null" json:"password" form:"password" valid:"required,minstringlength(8)~Password minimal 8 character"`
	Age      int     `gorm:"not null" json:"age" form:"age" valid:"required,minstringlength(8)~Age minimal 8 character"`
	Photos   []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	FullName string
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	u.Password = helpers.HashPassword(u.Password)
	return nil
}
