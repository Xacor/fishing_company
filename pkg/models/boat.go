package models

import (
	"time"

	"gorm.io/gorm"
)

type Boat struct {
	ID           int       `gorm:"primaryKey"`
	Name         string    `gorm:"size:124;unique;not null"`
	BtypeID      uint8     `gorm:"not null"`
	Displacement uint16    `gorm:"not null"`
	Build_date   time.Time `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Btype        Btype
}
