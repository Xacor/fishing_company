package db

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Xacor/fishing_company/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
			db, err := gorm.Open(dialecor)
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
	DB = db
}
