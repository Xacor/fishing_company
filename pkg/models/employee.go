package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	ID         int    `gorm:"primaryKey"`
	Lastname   string `gorm:"size:124;not null"`
	Firstname  string `gorm:"size:124;not null"`
	Middlename string `gorm:"size:124;"`
	Address    string `gorm:"size:255;not null"`
	Birth_date time.Time
	PositionID uint8 `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Position   Position
}

type Position struct {
	ID   uint8  `gorm:"primaryKey"`
	Name string `gorm:"size:124"`
}
