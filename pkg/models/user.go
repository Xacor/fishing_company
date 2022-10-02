package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"unique; size:32"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
