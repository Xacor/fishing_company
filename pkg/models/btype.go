package models

type Btype struct {
	ID   uint8  `gorm:"primaryKey"`
	Name string `gorm:"size:124"`
}
