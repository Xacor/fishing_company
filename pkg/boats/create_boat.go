package boats

import (
	"fishing_company/pkg/models"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h handler) BoatForm(c *gin.Context) {
	c.HTML(http.StatusOK, "create_boat.html", gin.H{})
}

func (h handler) CreateBoat(c *gin.Context) {
	var boat models.Boat
	boat.Name = c.PostForm("name")

	typeID, _ := strconv.Atoi(c.PostForm("type"))
	boat.BtypeID = uint8(typeID)

	displacement, _ := strconv.Atoi(c.PostForm("displacement"))
	boat.Displacement = uint16(displacement)

	//"2006-01-02" лучше вынести константой
	date, _ := time.Parse("2006-01-02", c.PostForm("build_date"))
	boat.Build_date = date

	if result := h.DB.Create(&boat); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	dest_url := url.URL{Path: fmt.Sprintf("/boats/%d", boat.ID)}
	c.Redirect(http.StatusMovedPermanently, dest_url.String())
}
