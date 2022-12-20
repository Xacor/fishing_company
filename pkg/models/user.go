package models

import (
	"time"
)

type User struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"unique; size:32"`
	Password  string
	RoleID    int `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Role      Role
}

type Role struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"size:16;unique;not null"`
}
