package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	TodoID      int64  `json:"id"`
	Title       string `gorm:"not null" json:"title" binding:"required"`
	Description string `gorm:"not null" json:"description" binding:"required"`
	CreatedBy   string `json:"-"`
}
