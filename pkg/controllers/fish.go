package controllers

import (
	"fishing_company/pkg/db"
	"fishing_company/pkg/models"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetFishes(c *gin.Context) {
	var fishes []models.FishType
	result := db.DB.Find(&fishes)
	if result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}
	c.HTML(http.StatusOK, "fishes", gin.H{
		"Number": result.RowsAffected,
		"fishes": &fishes,
	})
}

func FishForm(c *gin.Context) {
	c.HTML(http.StatusOK, "fishForm", gin.H{})
}

func CreateFish(c *gin.Context) {
	var fish models.FishType
	err := c.ShouldBind(&fish)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if result := db.DB.Create(&fish); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}
	dest_url := url.URL{Path: "/fishes"}
	c.Redirect(http.StatusMovedPermanently, dest_url.String())
}

func DeleteFish(c *gin.Context) {
	fishID := c.Param("id")
	var fish models.FishType
	//_ = db.DB.First(&fish, fishID)
	if result := db.DB.Delete(&fish, fishID); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}
	dest_url := url.URL{Path: "/fishes"}
	c.Redirect(http.StatusMovedPermanently, dest_url.String())
}
