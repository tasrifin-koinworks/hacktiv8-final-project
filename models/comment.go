package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserID  int `json:"user_id"`
	User    *User
	PhotoID int `json:"photo_id"`
	Photo   *Photo
	Message string `gorm:"not null;" json:"message" valid:"required~ Message is required"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	return errCreate
}
