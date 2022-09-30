package boats

import (
	"fishing_company/pkg/models"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteBoat(c *gin.Context) {
	boatId := c.Param("id")

	var boat models.Boat

	if result := h.DB.Delete(&boat, boatId); result.Error != nil {
		fmt.Println(result.Error)
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	dest_url := url.URL{Path: "/boats"}
	c.Redirect(http.StatusMovedPermanently, dest_url.String())

}
