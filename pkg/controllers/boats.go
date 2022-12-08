package controllers

import (
	"fishing_company/pkg/db"
	"fishing_company/pkg/globals"
	"fishing_company/pkg/models"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func BoatForm(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	c.HTML(http.StatusOK, "createBoat", gin.H{"user": user})
}

func CreateBoat(c *gin.Context) {
	var boat models.Boat
	boat.Name = c.PostForm("name")

	typeID, _ := strconv.Atoi(c.PostForm("type"))
	boat.BtypeID = uint8(typeID)

	displacement, _ := strconv.Atoi(c.PostForm("displacement"))
	boat.Displacement = uint16(displacement)

	//"2006-01-02" лучше вынести константой
	date, _ := time.Parse("2006-01-02", c.PostForm("build_date"))
	boat.Build_date = date

	if result := db.DB.Create(&boat); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return
	}
	dest_url := url.URL{Path: fmt.Sprintf("/boats/%d", boat.ID)}
	c.Redirect(http.StatusMovedPermanently, dest_url.String())
}

func DeleteBoat(c *gin.Context) {
	boatID := c.Param("id")
	if c.PostForm("boatName") != c.PostForm("inputBoatName") {
		c.Redirect(http.StatusMovedPermanently, "/boats")

		return
	}
	var boat models.Boat
	_ = db.DB.First(&boat, boatID)
	if result := db.DB.Delete(&boat); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Print(err)
		}

		return
	}

	dest_url := url.URL{Path: "/boats"}
	c.Redirect(http.StatusMovedPermanently, dest_url.String())

}

func GetBoat(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	var boat models.Boat

	id := c.Param("id")

	if result := db.DB.First(&boat, id); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return
	}
	c.HTML(http.StatusOK, "boat", gin.H{
		"boat": boat,
		"user": user,
	})
}

func GetBoats(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	var boats []models.Boat
	result := db.DB.Find(&boats)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return
	}
	c.HTML(http.StatusOK, "boats", gin.H{
		"Number": result.RowsAffected,
		"Boats":  &boats,
		"user":   user,
	})
}

func UpdateBoatForm(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	boatId := c.Param("id")

	var boat models.Boat

	if result := db.DB.First(&boat, boatId); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return
	}

	c.HTML(http.StatusOK, "updateBoat", gin.H{
		"boat": boat,
		"user": user,
	})
}

func UpdateBoat(c *gin.Context) {

	boatId := c.Param("id")

	var boat models.Boat

	if result := db.DB.First(&boat, boatId); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

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
		if err := c.AbortWithError(http.StatusNotModified, result.Error); err != nil {
			log.Println(err)
		}

		return
	}

	dest_url := url.URL{Path: fmt.Sprintf("/boats/%d", boat.ID)}
	c.Redirect(http.StatusMovedPermanently, dest_url.String())

}
