package boats

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"fishing_company/pkg/models"

	"github.com/gin-gonic/gin"
)

func (h handler) UpdateBoatForm(c *gin.Context) {
	boatId := c.Param("id")

	var boat models.Boat

	if result := h.DB.First(&boat, boatId); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.HTML(http.StatusOK, "update_boat.html", gin.H{
		"boat": boat,
	})
}

func (h handler) UpdateBoat(c *gin.Context) {

	boatId := c.Param("id")

	var boat models.Boat

	if result := h.DB.First(&boat, boatId); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
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

	if result := h.DB.Save(&boat); result.Error != nil {
		c.AbortWithError(http.StatusNotModified, result.Error)
		return
	}

	dest_url := url.URL{Path: fmt.Sprintf("/boats/%d", boat.ID)}
	c.Redirect(http.StatusMovedPermanently, dest_url.String())

}
