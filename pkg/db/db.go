package db

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Xacor/fishing_company/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectLoop(timeout time.Duration, dialecor gorm.Dialector) (db *gorm.DB, err error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	timeoutExceeded := time.After(timeout)
	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %s timeout", timeout)

		case <-ticker.C:
			db, err := gorm.Open(dialecor, &gorm.Config{
				Logger: logger.Default.LogMode(logger.Silent),
			})
			if err == nil {
				return db, nil
			}
			log.Error(err)
		}
	}
}

func Init(url string) {
	db, err := ConnectLoop(time.Second*30, mysql.Open(url))

	if err != nil {
		log.Fatalln(err)
	}

	if err = db.SetupJoinTable(&models.Trip{}, "FishTypes", &models.FishTypeTrip{}); err != nil {
		log.Fatalln(err)
	}

	if err = db.AutoMigrate(&models.Boat{}, &models.FishTypeTrip{}, &models.Trip{}, &models.User{}, &models.Employee{}, &models.Position{}, &models.FishType{}, &models.SeaBank{}, &models.Btype{}); err != nil {
		log.Fatalln(err)
	}

	result := db.Find(&[]models.Role{})
	if result.RowsAffected == 0 {
		result := db.Create(&models.Role{ID: 1, Name: "Admin"}).Create(&models.Role{ID: 2, Name: "User"})
		if result.Error != nil {
			log.Fatal(result.Error)
		}
	}

	DB = db

}
