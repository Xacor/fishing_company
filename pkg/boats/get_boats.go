package boats

import (
	"log"
	"net/http"

	"fishing_company/pkg/db"
	"fishing_company/pkg/models"

	"github.com/gin-gonic/gin"
)

func GetBoats(c *gin.Context) {
	var boats []models.Boat
	result := db.DB.Find(&boats)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return
	}
	c.HTML(http.StatusOK, "all_boats.html", gin.H{
		"Number": result.RowsAffected,
		"Boats":  &boats,
	})
}
