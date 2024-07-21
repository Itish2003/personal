package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title   string `gorm:"unique" json:"title"`
	Content string `json:"content"`
}
