package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Xacor/fishing_company/pkg/db"
	"github.com/Xacor/fishing_company/pkg/globals"
	"github.com/Xacor/fishing_company/pkg/models"
	"github.com/Xacor/fishing_company/pkg/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func BoatForm(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	var boatTypes []models.Btype

	if err := db.DB.Find(&boatTypes).Error; err != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "createBoat", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	c.HTML(http.StatusOK, "createBoat", gin.H{
		"user":      user,
		"alerts":    utils.Flashes(c),
		"boatTypes": boatTypes,
	})
}

func CreateBoat(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	var boat models.Boat
	boat.Name = c.PostForm("name")

	typeID, _ := strconv.Atoi(c.PostForm("type"))
	boat.BtypeID = uint8(typeID)

	displacement, _ := strconv.Atoi(c.PostForm("displacement"))
	boat.Displacement = uint16(displacement)

	date, _ := time.Parse(globals.TimeLayout, c.PostForm("build_date"))
	boat.Build_date = date

	if err := db.DB.Where("name = ?", boat.Name).First(&models.Boat{}).Error; err == nil {
		utils.FlashMessage(c, "Судно с таким именем уже существует или существовало")
		c.HTML(http.StatusBadRequest, "createBoat", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	if result := db.DB.Create(&boat); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "boats", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}
	dest_url := url.URL{Path: fmt.Sprintf("/boats/%d", boat.ID)}
	c.Redirect(http.StatusSeeOther, dest_url.String())
}

func DeleteBoat(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	boatID := c.Param("id")

	if result := db.DB.Delete(&models.Boat{}, boatID); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "boats", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	dest_url := url.URL{Path: "/boats"}
	c.Redirect(http.StatusSeeOther, dest_url.String())

}

func GetBoat(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	var boat models.Boat
	id := c.Param("id")

	if result := db.DB.First(&boat, id); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "boat", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}
	c.HTML(http.StatusOK, "boat", gin.H{
		"boat":   boat,
		"user":   user,
		"alerts": utils.Flashes(c),
	})
}

func GetBoats(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	var boats []models.Boat
	result := db.DB.Find(&boats)
	if result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "boats", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}
	c.HTML(http.StatusOK, "boats", gin.H{
		"Number": result.RowsAffected,
		"Boats":  &boats,
		"user":   user,
		"alerts": utils.Flashes(c),
	})
}

func UpdateBoatForm(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	boatId := c.Param("id")
	var boat models.Boat

	var boatTypes []models.Btype

	if err := db.DB.Find(&boatTypes).Error; err != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "updateBoat", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	if result := db.DB.First(&boat, boatId); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "updateBoat", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	c.HTML(http.StatusOK, "updateBoat", gin.H{
		"boat":      boat,
		"user":      user,
		"boatTypes": boatTypes,
	})
}

func UpdateBoat(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	boatId := c.Param("id")
	var boat models.Boat

	if result := db.DB.First(&boat, boatId); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "updateBoat", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	if boatName := c.PostForm("name"); boatName != "" {
		boat.Name = boatName
	}

	if c.PostForm("type") != "" {
		boatTypeId, _ := strconv.Atoi(c.PostForm("type"))
		boat.BtypeID = uint8(boatTypeId)
	}

	if c.PostForm("displacement") != "" {
		boatDisplacement, _ := strconv.Atoi(c.PostForm("displacement"))
		boat.Displacement = uint16(boatDisplacement)
	}

	if result := db.DB.Save(&boat); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "updateBoat", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	dest_url := url.URL{Path: fmt.Sprintf("/boats/%d", boat.ID)}
	c.Redirect(http.StatusSeeOther, dest_url.String())

}
