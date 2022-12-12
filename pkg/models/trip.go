package models

import "time"

type Trip struct {
	ID            int `gorm:"primaryKey"`
	BoatID        int `gorm:"not null"`
	DepartureDate time.Time
	ArrivalDate   time.Time `gorm:"default:2006-01-02"`
	Boat          Boat
	Employees     []Employee `gorm:"many2many:trips_employees;"`
	SeaBanks      []SeaBank  `gorm:"many2many:sea_banks_trips;"`
	FishTypes     []FishType `gorm:"many2many:fish_type_trips;"`
}

type FishTypeTrip struct {
	FishTypeID int
	TripID     int
	Catch      int `gorm:"default:0"`
}
