package controllers

import (
	"fishing_company/pkg/db"
	"fishing_company/pkg/globals"
	"fishing_company/pkg/models"
	"fishing_company/pkg/utils"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func EmployeeForm(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	c.HTML(http.StatusOK, "createEmployee", gin.H{
		"user":   user,
		"alerts": utils.Flashes(c),
	})
}

func CreateEmployee(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	var employee models.Employee

	employee.Lastname = c.PostForm("lastname")
	employee.Firstname = c.PostForm("firstname")
	employee.Middlename = c.PostForm("middlename")
	employee.Address = c.PostForm("address")

	date, _ := time.Parse(globals.TimeLayout, c.PostForm("birth_date"))
	employee.Birth_date = date

	positionID, _ := strconv.Atoi(c.PostForm("position"))
	employee.PositionID = uint8(positionID)

	if result := db.DB.Create(&employee); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "bankForm", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}
	dest_url := url.URL{Path: fmt.Sprintf("/employees/%d", employee.ID)}
	c.Redirect(http.StatusSeeOther, dest_url.String())
}

func DeleteEmployee(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	employeeID := c.Param("id")
	var employee models.Employee

	if result := db.DB.Delete(&employee, employeeID); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "employees", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	dest_url := url.URL{Path: "/employees"}
	c.Redirect(http.StatusSeeOther, dest_url.String())

}

func GetEmployee(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	var employee models.Employee
	id := c.Param("id")

	if result := db.DB.First(&employee, id); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "employee", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}
	c.HTML(http.StatusOK, "employee", gin.H{
		"employee": employee,
		"user":     user,
	})
}

func GetEmployees(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	var employees []models.Employee
	result := db.DB.Find(&employees)
	if result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "employees", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}
	c.HTML(http.StatusOK, "employees", gin.H{
		"Number":    result.RowsAffected,
		"Employees": &employees,
		"user":      user,
	})
}

func UpdateEmployeeForm(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	employeeId := c.Param("id")
	var employee models.Employee

	if result := db.DB.First(&employee, employeeId); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "updateEmployee", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	c.HTML(http.StatusOK, "updateEmployee", gin.H{
		"employee": employee,
		"user":     user,
	})
}

func UpdateEmployee(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	employeeId := c.Param("id")

	var employee models.Employee

	if result := db.DB.First(&employee, employeeId); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "updateEmployee", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	if lastname := c.PostForm("lastname"); lastname != "" {
		employee.Lastname = lastname
	}
	if firstname := c.PostForm("firstname"); firstname != "" {
		employee.Firstname = firstname
	}
	if middlename := c.PostForm("middlename"); middlename != "" {
		employee.Firstname = middlename
	}
	if address := c.PostForm("address"); address != "" {
		employee.Firstname = address
	}
	if c.PostForm("birth_date") != "" {
		date, _ := time.Parse(globals.TimeLayout, c.PostForm("birth_date"))
		employee.Birth_date = date
	}

	if c.PostForm("position") != "" {
		positionID, _ := strconv.Atoi(c.PostForm("position"))
		employee.PositionID = uint8(positionID)
	}

	if result := db.DB.Save(&employee); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "updateEmployee", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	dest_url := url.URL{Path: fmt.Sprintf("/employees/%d", employee.ID)}
	c.Redirect(http.StatusSeeOther, dest_url.String())

}
