package boats

import (
	"net/http"

	"fishing_company/pkg/db"
	"fishing_company/pkg/models"

	"github.com/gin-gonic/gin"
)

func GetBoat(c *gin.Context) {
	var boat models.Boat

	id := c.Param("id")

	if result := db.DB.First(&boat, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.HTML(http.StatusOK, "show_boat.html", gin.H{
		"boat": boat,
	})
}
