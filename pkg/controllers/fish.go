package controllers

import (
	"fishing_company/pkg/db"
	"fishing_company/pkg/globals"
	"fishing_company/pkg/models"
	"fishing_company/pkg/utils"
	"net/http"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetFishes(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	var fishes []models.FishType
	result := db.DB.Find(&fishes)
	if result.Error != nil {
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
		utils.FlashMessage(c, "Возникла ошибка при обработке формы")
		c.HTML(http.StatusBadRequest, "fishForm", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	if result := db.DB.Create(&fish); result.Error != nil {
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
