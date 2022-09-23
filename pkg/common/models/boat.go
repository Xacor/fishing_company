package models

import (
	"gorm.io/datatypes"
)

type Boat struct {
	ID           int            `gorm:"primaryKey" json:"id"`
	Name         string         `json:"name"`
	Type         int            `json:"type"`
	Displacement int            `json:"displacement"`
	Build_date   datatypes.Date `json:"build_date"`
}
