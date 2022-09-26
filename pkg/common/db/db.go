package db

import (
	"fmt"
	"log"
	"time"

	"fishing_company/pkg/common/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectLoop(timeout time.Duration, dialecor gorm.Dialector) (db *gorm.DB, err error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeoutExceeded := time.After(timeout)
	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %s timeout", timeout)

		case <-ticker.C:
			db, err := gorm.Open(dialecor, &gorm.Config{})
			if err == nil {
				return db, nil
			}
			log.Println(err)
		}
	}
}

func Init(url string) *gorm.DB {
	db, err := ConnectLoop(time.Second*10, mysql.Open(url))

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Boat{})
	return db
}
