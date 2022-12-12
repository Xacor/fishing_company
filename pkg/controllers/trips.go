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

func getFreeBoats(c *gin.Context) []models.Boat {
	var bIDs []int
	var boats []models.Boat
	// select boat_id from trips where arrival_date = '2006-01-02';
	result := db.DB.Model(&models.Trip{}).Select("boat_id").Where("arrival_date = ?", "2006-01-02").Find(&bIDs)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return nil
	}

	result = db.DB.Not(bIDs).Find(&boats)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return nil
	}
	return boats

}

func getFreeEmployees(c *gin.Context) []models.Employee {
	type TripsEmployees struct {
		TripID     int
		EmployeeID int
	}
	var empIDs []int
	var employees []models.Employee
	//select employee_id from trips_employees inner join trips on trip_id = trips.id where trips.arrival_date = '2006-01-02';
	result := db.DB.Model(&TripsEmployees{}).Select("employee_id").Joins("inner join trips on trip_id = trips.id").Where("trips.arrival_date = ?", "2006-01-02").Find(&empIDs)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return nil
	}

	result = db.DB.Not(empIDs).Find(&employees)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}

		return nil
	}

	return employees
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

	boats = getFreeBoats(c)

	employees = getFreeEmployees(c)

	sbs := getFormSeaBanks(c)
	fts := getFormFishTypes(c)

	c.HTML(http.StatusOK, "createTrip", gin.H{
		"Boats":     &boats,
		"Employees": &employees,
		"SeaBanks":  &sbs,
		"FishTypes": &fts,
	})
}

func convertStrSliceToIntSlice(slc []string) []int {
	l := len(slc)
	if l != 0 {
		slice := make([]int, l)
		for i, id := range slc {
			v, _ := strconv.Atoi(id)
			slice[i] = int(v)
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
	emps := convertStrSliceToIntSlice(empStrIDs)
	getEmployees(c, &trip, emps)

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
	result := db.DB.Preload("Boat").Preload("Employees").Preload("FishTypes").Preload("SeaBanks").Find(&trips)
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
	result := db.DB.Preload("Boat").Preload("Employees").Preload("FishTypes").Preload("SeaBanks").First(&trip, id)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.HTML(http.StatusOK, "trip", gin.H{"Trip": &trip})
}
