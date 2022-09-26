package models

import (
	"time"
)

type Boat struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	Name         string    `json:"name"`
	Type         int       `json:"type"`
	Displacement int       `json:"displacement"`
	Build_date   time.Time `json:"build_date"`
}
