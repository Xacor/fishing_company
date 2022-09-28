package models

import (
	"time"
)

type Boat struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Type         int
	Displacement int
	Build_date   time.Time
}
