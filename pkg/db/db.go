package db

import (
	"fmt"
	"log"
	"time"

	"fishing_company/pkg/models"

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
			log.Println(err)
		}
	}
}

func Init(url string) {
	db, err := ConnectLoop(time.Second*10, mysql.Open(url))

	if err != nil {
		log.Fatalln(err)
	}

	if err = db.AutoMigrate(&models.Boat{}, &models.User{}); err != nil {
		log.Fatalln(err)
	}
	DB = db
}
