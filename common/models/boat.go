package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Boat struct {
	gorm.Model
	ID           int            `gorm:"primaryKey" json:"id"`
	Name         string         `json:"name"`
	Type         int            `json:"type"`
	Displacement int            `json:"displacement"`
	Build_date   datatypes.Date `json:"build_date"`
}
