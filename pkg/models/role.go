package models

type Role struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"size:16;unique;not null"`
}
