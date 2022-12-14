package controllers

import (
	"net/http"
	"net/url"

	"github.com/Xacor/fishing_company/pkg/db"
	"github.com/Xacor/fishing_company/pkg/globals"
	"github.com/Xacor/fishing_company/pkg/models"
	"github.com/Xacor/fishing_company/pkg/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetFishes(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	var fishes []models.FishType
	result := db.DB.Find(&fishes)
	if result.Error != nil {
		log.Error(result.Error)
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "fishes", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	c.HTML(http.StatusOK, "fishes", gin.H{
		"Number": result.RowsAffected,
		"fishes": &fishes,
		"user":   user,
		"alerts": utils.Flashes(c),
	})
}

func FishForm(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	c.HTML(http.StatusOK, "fishForm", gin.H{
		"user":   user,
		"alerts": utils.Flashes(c),
	})
}

func CreateFish(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	var fish models.FishType

	if err := c.ShouldBind(&fish); err != nil {
		log.Error(err)
		utils.FlashMessage(c, "Возникла ошибка при обработке формы")
		c.HTML(http.StatusBadRequest, "fishForm", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	if result := db.DB.Create(&fish); result.Error != nil {
		log.Error(result.Error)
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "fishForm", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}
	dest_url := url.URL{Path: "/fishes"}
	c.Redirect(http.StatusSeeOther, dest_url.String())
}

func DeleteFish(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	fishID := c.Param("id")
	var fish models.FishType

	if result := db.DB.Delete(&fish, fishID); result.Error != nil {
		log.Error(result.Error)
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "fishForm", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}
	dest_url := url.URL{Path: "/fishes"}
	c.Redirect(http.StatusSeeOther, dest_url.String())
}
