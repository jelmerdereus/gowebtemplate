package models

import (
	"github.com/jinzhu/gorm"
)

// Event contains operational API and database events
type Event struct {
	gorm.Model
	UserID  uint   `gorm:"column:user_id" json:"USER_ID"`
	Action  string `gorm:"column:action" json:"Action"`
	Type    string `gorm:"column:type" json:"Type"`
	Message string `gorm:"column:message" json:"Message"`
}
