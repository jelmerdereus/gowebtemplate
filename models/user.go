package models

import (
	"github.com/jinzhu/gorm"
)

// User is a user object
type User struct {
	gorm.Model
	Alias    string `gorm:"column:alias;unique" json:"Alias"`
	Password string `gorm:"column:password" json:"PasswordHash"`
}
