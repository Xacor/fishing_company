package boats

import (
	"log"
	"net/http"

	"fishing_company/pkg/db"
	"fishing_company/pkg/models"

	"github.com/gin-gonic/gin"
)

func GetBoat(c *gin.Context) {
	var boat models.Boat

	id := c.Param("id")

	if result := db.DB.First(&boat, id); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return
	}
	c.HTML(http.StatusOK, "show_boat.html", gin.H{
		"boat": boat,
	})
}
