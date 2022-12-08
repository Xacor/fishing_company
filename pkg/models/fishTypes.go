package models

type FishType struct {
	ID       int       `gorm:"primaryKey"`
	Name     string    `gorm:"not null" form:"name"`
	SeaBanks []SeaBank `gorm:"many2many:sea_banks_fish_types;"`
	Trips    []Trip    `gorm:"many2many:fish_types_trips;"`
}
