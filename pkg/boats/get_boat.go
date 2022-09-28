package boats

import (
	"net/http"

	"fishing_company/pkg/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetBoat(c *gin.Context) {
	var boat models.Boat

	id := c.Param("id")

	if result := h.DB.First(&boat, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.HTML(http.StatusOK, "show_boat.html", gin.H{
		"boat": boat,
	})
}
