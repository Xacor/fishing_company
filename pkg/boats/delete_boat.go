package boats

import (
	"fishing_company/pkg/db"
	"fishing_company/pkg/models"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func DeleteBoatForm(c *gin.Context) {
	boatID := c.Param("id")
	var boat models.Boat

	if result := db.DB.First(&boat, boatID); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.HTML(http.StatusOK, "delete_boat.html", gin.H{
		"boatID":   boatID,
		"boatName": boat.Name,
	})
}

func DeleteBoat(c *gin.Context) {
	boatID := c.Param("id")
	if c.PostForm("boatName") != c.PostForm("inputBoatName") {
		c.Redirect(http.StatusMovedPermanently, "/boats")
		return
	}
	var boat models.Boat
	result := db.DB.First(&boat, boatID)
	if result = db.DB.Delete(&boat); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	dest_url := url.URL{Path: "/boats"}
	c.Redirect(http.StatusMovedPermanently, dest_url.String())

}
