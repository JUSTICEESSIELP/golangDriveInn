package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID             uint      `gorm:"primaryKey"`
	FirstName      string    `json"firstName`
	LastName       string    `json"lastName"`
	Email          string    `json"email"`
	Password       string    `json"password"`
	SocialLogin    bool      `json"socialLogin"`
	SocialProvider string    `json"socialProvider"`
	CreatedAt      time.Time `json"createdAt"`
	UpdatedAt      time.Time `json"updatedAt"`
}
