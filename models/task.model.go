package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `json:"title" gorm:"text;not null"`
	Description string `json:"description" gorm:"text;default:null"`
	Done        bool   `json:"done" gorm:"default:false"`
}
