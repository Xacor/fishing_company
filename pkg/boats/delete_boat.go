package boats

import (
	"fishing_company/pkg/models"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteBoatForm(c *gin.Context) {
	boatID := c.Param("id")
	var boat models.Boat

	if result := h.DB.First(&boat, boatID); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.HTML(http.StatusOK, "delete_boat.html", gin.H{
		"boatID":   boatID,
		"boatName": boat.Name,
	})
}

func (h handler) DeleteBoat(c *gin.Context) {
	boatID := c.Param("id")
	if c.PostForm("boatName") != c.PostForm("inputBoatName") {
		c.Redirect(http.StatusMovedPermanently, "/boats")
		return
	}
	var boat models.Boat
	result := h.DB.First(&boat, boatID)
	if result = h.DB.Delete(&boat); result.Error != nil {
		fmt.Println("AAAAAAAAAAAAAAAA")
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	dest_url := url.URL{Path: "/boats"}
	c.Redirect(http.StatusMovedPermanently, dest_url.String())

}
