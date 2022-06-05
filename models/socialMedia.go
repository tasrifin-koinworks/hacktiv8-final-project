package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null;" json:"name" valid:"required~ Your name is required"`
	SocialMediaUrl string `gorm:"not null;" json:"social_media_url" valid:"required~ Your social media url is required"`
	UserID         int
	User           *User
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)
	return errCreate
}
