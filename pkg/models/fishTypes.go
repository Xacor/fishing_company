package models

type FishType struct {
	ID       int       `gorm:"primaryKey"`
	Name     string    `gorm:"not null"`
	SeaBanks []SeaBank `gorm:"many2many:sea_banks_fish_types;"`
}
