package boats

import (
	"fishing_company/pkg/common/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

func (h handler) BoatForm(c *gin.Context) {
	c.HTML(http.StatusOK, "create_boat.html", gin.H{})
}

func (h handler) CreateBoat(c *gin.Context) {
	var boat models.Boat
	boat.Name = c.PostForm("name")
	boat.Type, _ = strconv.Atoi(c.PostForm("type"))
	boat.Displacement, _ = strconv.Atoi(c.PostForm("displacement"))

	//"02/01/2006" лучше вынести константой
	date, _ := time.Parse("2006-01-02", c.PostForm("build_date"))
	boat.Build_date = datatypes.Date(date)

	if result := h.DB.Create(&boat); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	//редирект в таком виде не работает
	c.Redirect(http.StatusCreated, "/")
}
