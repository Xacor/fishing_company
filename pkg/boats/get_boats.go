package boats

import (
	"net/http"

	"fishing_company/pkg/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetBoats(c *gin.Context) {
	var boats []models.Boat
	result := h.DB.Find(&boats)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.HTML(http.StatusOK, "all_boats.html", gin.H{
		"Number": result.RowsAffected,
		"Boats":  &boats,
	})
}
