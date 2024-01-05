package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	SocialLogin    bool   `json:"social_login"`
	SocialProvider string `json:"social_provider"`

}
