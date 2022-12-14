package utils

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/Xacor/fishing_company/pkg/db"
	"github.com/Xacor/fishing_company/pkg/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Returns user if found or error if not
func CheckUserPass(username, password string) (*models.User, error) {

	var user models.User
	if err := db.DB.Preload("Role").Where(&models.User{Name: username}).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return &user, nil

}

func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}

func FlashMessage(c *gin.Context, message string) {
	session := sessions.Default(c)
	session.AddFlash(message)
	if err := session.Save(); err != nil {
		log.Error(err)
	}
}

func Flashes(c *gin.Context) []interface{} {
	session := sessions.Default(c)
	flashes := session.Flashes()
	if len(flashes) != 0 {
		if err := session.Save(); err != nil {
			log.Error(err)
		}
	}
	return flashes
}
