package boats

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"fishing_company/pkg/db"
	"fishing_company/pkg/models"

	"github.com/gin-gonic/gin"
)

func UpdateBoatForm(c *gin.Context) {
	boatId := c.Param("id")

	var boat models.Boat

	if result := db.DB.First(&boat, boatId); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return
	}

	c.HTML(http.StatusOK, "update_boat.html", gin.H{
		"boat": boat,
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
