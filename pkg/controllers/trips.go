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

func getBoatsID(c *gin.Context) []int {
	type BoatID struct {
		BoatID int
	}
	var bIDs []BoatID
	result := db.DB.Model(&models.Trip{}).Where("arrival_date = ?", "2006-01-02").Find(&bIDs)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return nil
	}
	listID := make([]int, 0, len(bIDs))
	for i, id := range bIDs {
		listID[i] = id.BoatID
	}
	log.Println(listID)

	return listID
}

func getEmployeesID(c *gin.Context) []int {
	type TripsEmployees struct {
		TripID     int
		EmployeeID int
	}

	var emps []models.Employee
	// result := db.DB.Model(&models.Employee{}).Preload("Trips").Where("arrival_date = ?", "2006-01-02").Find(&emps)
	// SELECT * FROM `employees` WHERE arrival_date = '2006-01-02' AND `employees`.`deleted_at` IS NULL

	// db.DB.Model(&TripsEmployees{}).Select("employee_id").Joins("inner join trips on trip_id = trips.id").Where("trips.arrival_date = ?", "2006-01-02").Scan(&res)
	result := db.DB.Model(&TripsEmployees{}).Select("employee_id").Joins("inner join trips on trip_id = trips.id").Where("trips.arrival_date = ?", "2006-01-02").Find(&emps)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return nil
	}
	listID := make([]int, 0, len(emps))
	for i, emp := range emps {
		listID[i] = emp.ID
	}
	return listID
}

func getEmployees(c *gin.Context, trip *models.Trip, empIntIDs []int) {
	result := db.DB.Find(&trip.Employees, empIntIDs)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}
	}
}

func getFormSeaBanks(c *gin.Context) []models.SeaBank {
	var sbs []models.SeaBank
	result := db.DB.Find(&sbs)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}
	}
	return sbs
}

func getFormFishTypes(c *gin.Context) []models.FishType {
	var fts []models.FishType
	result := db.DB.Find(&fts)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}
	}
	return fts
}

func TripForm(c *gin.Context) {
	var boats []models.Boat
	var employees []models.Employee

	listID := getBoatsID(c)
	result := db.DB.Not(listID).Find(&boats)

	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return
	}

	listID = getEmployeesID(c)
	result = db.DB.Not(listID).Find(&employees)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return
	}

	sbs := getFormSeaBanks(c)
	fts := getFormFishTypes(c)

	c.HTML(http.StatusOK, "createTrip", gin.H{
		"Boats":     &boats,
		"Employees": &employees,
		"SeaBanks":  &sbs,
		"FishTypes": &fts,
	})
}

// какая-то ошибка здесь вылезает
func convertStrSliceToIntSlice(slc []string) []int {
	if len(slc) != 0 {
		slice := make([]int, 0, len(slc))
		for i, id := range slc {
			slice[i], _ = strconv.Atoi(id)
		}
		return slice
	}
	return nil
}

func getSeaBanks(c *gin.Context, trip *models.Trip, ids []int) {
	var sbs []models.SeaBank
	result := db.DB.Find(&sbs, ids)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}
	}
	trip.SeaBanks = append(trip.SeaBanks, sbs...)
}

func getFishTypes(c *gin.Context, trip *models.Trip, ids []int) {
	var fts []models.FishType
	result := db.DB.Find(&fts, ids)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}
	}

	trip.FishTypes = append(trip.FishTypes, fts...)
}

func CreateTrip(c *gin.Context) {
	var trip models.Trip

	boatID, _ := strconv.Atoi(c.PostForm("boats"))
	trip.BoatID = int(boatID)

	trip.DepartureDate = time.Now()
	trip.ArrivalDate = time.Time{}

	empStrIDs := c.PostFormArray("employees")
	log.Println(empStrIDs)
	// emps := convertStrSliceToIntSlice(empStrIDs)
	// getEmployees(c, &trip, emps)

	sbsStr := c.PostFormArray("seabanks")
	log.Println(sbsStr)
	sbs := convertStrSliceToIntSlice(sbsStr)
	getSeaBanks(c, &trip, sbs)

	ftsStr := c.PostFormArray("fish")
	log.Println(ftsStr)
	fts := convertStrSliceToIntSlice(ftsStr)
	getFishTypes(c, &trip, fts)

	if result := db.DB.Create(&trip); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return
	}
	dest_url := url.URL{Path: fmt.Sprintf("/trips/%d", trip.ID)}
	c.Redirect(http.StatusMovedPermanently, dest_url.String())
}

func GetTrips(c *gin.Context) {
	var trips []models.Trip
	result := db.DB.Find(&trips)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.HTML(http.StatusOK, "trips", gin.H{
		"Number": result.RowsAffected,
		"trips":  &trips,
	})
}

func GetTrip(c *gin.Context) {
	var trip models.Trip
	id := c.Param("id")

	if result := db.DB.First(&trip, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.HTML(http.StatusOK, "trip", gin.H{"Bank": &trip})
}
