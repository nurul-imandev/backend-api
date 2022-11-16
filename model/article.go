package model

import "time"

type Article struct {
	ID          uint   `gorm:"autoIncrement;not null"`
	Title       string `gorm:"size:100;not null"`
	Description string `gorm:"type:text;not null"`
	User        User
	UserID      uint `gorm:"index;not null"`
	Category    Category
	CategoryID  uint `gorm:"index;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
