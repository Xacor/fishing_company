package controllers

import (
	"fishing_company/pkg/db"
	"fishing_company/pkg/models"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetBank(c *gin.Context) {
	var bank models.SeaBank
	id := c.Param("id")

	if result := db.DB.First(&bank, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.HTML(http.StatusOK, "bank", gin.H{"Bank": &bank})
}

func GetBanks(c *gin.Context) {
	var banks []models.SeaBank
	result := db.DB.Find(&banks)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.HTML(http.StatusOK, "banks", gin.H{
		"Number": result.RowsAffected,
		"banks":  &banks,
	})
}

func BankForm(c *gin.Context) {
	c.HTML(http.StatusOK, "bankForm", gin.H{})
}

// нужно добавить валидацию
func CreateBank(c *gin.Context) {
	var bank models.SeaBank
	err := c.ShouldBind(&bank)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if result := db.DB.Create(&bank); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	dest_url := url.URL{Path: fmt.Sprintf("/banks/%d", bank.ID)}
	c.Redirect(http.StatusMovedPermanently, dest_url.String())
}

func DeleteBank(c *gin.Context) {
	bankID := c.Param("id")
	var bank models.SeaBank
	_ = db.DB.First(&bank, bankID)
	if result := db.DB.Delete(&bank); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	dest_url := url.URL{Path: "/banks"}
	c.Redirect(http.StatusMovedPermanently, dest_url.String())
}
