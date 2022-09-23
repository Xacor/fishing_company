package boats

import (
	"net/http"

	"fishing_company/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetBoat(c *gin.Context) {
	var boat models.Boat
	if result := h.DB.First(&boat); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.HTML(http.StatusOK, "base.html", gin.H{
		"Name": &boat.Name,
	})
}
