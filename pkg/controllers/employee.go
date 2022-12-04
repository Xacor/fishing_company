package controllers

import (
	"fishing_company/pkg/db"
	"fishing_company/pkg/models"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func EmployeeForm(c *gin.Context) {
	c.HTML(http.StatusOK, "createEmployee", gin.H{})
}

func CreateEmployee(c *gin.Context) {
	var employee models.Employee

	employee.Lastname = c.PostForm("lastname")
	employee.Firstname = c.PostForm("firstname")
	employee.Middlename = c.PostForm("middlename")
	employee.Address = c.PostForm("address")

	date, _ := time.Parse("2006-01-02", c.PostForm("birth_date"))
	employee.Birth_date = date

	positionID, _ := strconv.Atoi(c.PostForm("position"))
	employee.PositionID = uint8(positionID)

	if result := db.DB.Create(&employee); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}
		return
	}
	dest_url := url.URL{Path: fmt.Sprintf("/employees/%d", employee.ID)}
	c.Redirect(http.StatusMovedPermanently, dest_url.String())
}

func DeleteEmployeeForm(c *gin.Context) {
	employeeID := c.Param("id")
	var employee models.Employee

	if result := db.DB.First(&employee, employeeID); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return
	}

	c.HTML(http.StatusOK, "deleteEmployee", gin.H{
		"employeeID":   employeeID,
		"employeeName": employee.Lastname + " " + employee.Firstname,
	})
}

func DeleteEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	if c.PostForm("employeeName") != c.PostForm("inputEmployeeName") {
		c.Redirect(http.StatusMovedPermanently, "/employees")

		return
	}
	var employee models.Employee
	_ = db.DB.First(&employee, employeeID)
	if result := db.DB.Delete(&employee); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Print(err)
		}

		return
	}

	dest_url := url.URL{Path: "/employees"}
	c.Redirect(http.StatusMovedPermanently, dest_url.String())

}

func GetEmployee(c *gin.Context) {
	var employee models.Employee

	id := c.Param("id")

	if result := db.DB.First(&employee, id); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return
	}
	c.HTML(http.StatusOK, "employee", gin.H{
		"employee": employee,
	})
}

func GetEmployees(c *gin.Context) {
	var employees []models.Employee
	result := db.DB.Find(&employees)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return
	}
	c.HTML(http.StatusOK, "employees", gin.H{
		"Number":    result.RowsAffected,
		"Employees": &employees,
	})
}

func UpdateEmployeeForm(c *gin.Context) {
	employeeId := c.Param("id")

	var employee models.Employee

	if result := db.DB.First(&employee, employeeId); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return
	}

	c.HTML(http.StatusOK, "updateEmployee", gin.H{
		"employee": employee,
	})
}

func UpdateEmployee(c *gin.Context) {

	employeeId := c.Param("id")

	var employee models.Employee

	if result := db.DB.First(&employee, employeeId); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

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
		date, _ := time.Parse("2006-01-02", c.PostForm("birth_date"))
		employee.Birth_date = date
	}

	if c.PostForm("position") != "" {
		positionID, _ := strconv.Atoi(c.PostForm("position"))
		employee.PositionID = uint8(positionID)
	}

	if result := db.DB.Save(&employee); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotModified, result.Error); err != nil {
			log.Println(err)
		}

		return
	}

	dest_url := url.URL{Path: fmt.Sprintf("/employees/%d", employee.ID)}
	c.Redirect(http.StatusMovedPermanently, dest_url.String())

}
