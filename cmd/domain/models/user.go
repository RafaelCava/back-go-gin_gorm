package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time `gorm:"index"`
}
