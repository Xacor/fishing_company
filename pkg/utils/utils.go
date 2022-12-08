package utils

import (
	"fishing_company/pkg/db"
	"fishing_company/pkg/models"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Returns user if found or error if not
func CheckUserPass(username, password string) (*models.User, error) {

	var user models.User
	if err := db.DB.Preload("Role").Where(&models.User{Name: username}).First(&user).Error; err != nil {
		return nil, err
	}

	log.Printf("%+v", user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return &user, nil

}

func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}
