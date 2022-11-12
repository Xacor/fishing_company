package models

type SeaBank struct {
	ID        int        `gorm:"primaryKey"`
	Lat       float64    `gorm:"not null" form:"lat"`
	Lng       float64    `gorm:"not null" form:"lng"`
	Name      string     `gorm:"not null" form:"name"`
	FishTypes []FishType `gorm:"many2many:sea_banks_fish_types;"`
}
