package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/Xacor/fishing_company/pkg/db"
	"github.com/Xacor/fishing_company/pkg/globals"
	"github.com/Xacor/fishing_company/pkg/models"
	"github.com/Xacor/fishing_company/pkg/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetBank(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	var bank models.SeaBank
	id := c.Param("id")

	if result := db.DB.First(&bank, id); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "bank", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	c.HTML(http.StatusOK, "bank", gin.H{
		"Bank": &bank,
		"user": user,
	})
}

func GetBanks(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	var banks []models.SeaBank
	result := db.DB.Find(&banks)
	if result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "banks", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	c.HTML(http.StatusOK, "banks", gin.H{
		"Number": result.RowsAffected,
		"banks":  &banks,
		"user":   user,
		"alerts": utils.Flashes(c),
	})
}

func BankForm(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	c.HTML(http.StatusOK, "bankForm", gin.H{
		"user":   user,
		"alerts": utils.Flashes(c),
	})
}

// нужно добавить валидацию
func CreateBank(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	var bank models.SeaBank
	err := c.ShouldBind(&bank)
	if err != nil {
		utils.FlashMessage(c, "Возникла ошибка при обработке формы")
		c.HTML(http.StatusBadRequest, "bankForm", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}
	if result := db.DB.Create(&bank); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "bankForm", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	dest_url := url.URL{Path: fmt.Sprintf("/banks/%d", bank.ID)}
	c.Redirect(http.StatusSeeOther, dest_url.String())
}

func DeleteBank(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	bankID := c.Param("id")
	var bank models.SeaBank
	_ = db.DB.First(&bank, bankID)
	if result := db.DB.Delete(&bank); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "banks", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	dest_url := url.URL{Path: "/banks"}
	c.Redirect(http.StatusSeeOther, dest_url.String())
}
