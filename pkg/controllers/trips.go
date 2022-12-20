package controllers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Xacor/fishing_company/pkg/db"
	"github.com/Xacor/fishing_company/pkg/globals"
	"github.com/Xacor/fishing_company/pkg/models"
	"github.com/Xacor/fishing_company/pkg/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetFreeBoats(c *gin.Context) []models.Boat {
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

func GetFreeEmployees(c *gin.Context) []models.Employee {
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

func GetEmployeesForTrip(c *gin.Context, trip *models.Trip, empIntIDs []int) {
	result := db.DB.Find(&trip.Employees, empIntIDs)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}
	}
}

func GetFormSeaBanks(c *gin.Context) []models.SeaBank {
	var sbs []models.SeaBank
	result := db.DB.Find(&sbs)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}
	}
	return sbs
}

func GetFormFishTypes(c *gin.Context) []models.FishType {
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

	boats = GetFreeBoats(c)

	employees = GetFreeEmployees(c)

	sbs := GetFormSeaBanks(c)
	fts := GetFormFishTypes(c)

	c.HTML(http.StatusOK, "createTrip", gin.H{
		"Boats":     &boats,
		"Employees": &employees,
		"SeaBanks":  &sbs,
		"FishTypes": &fts,
	})
}

func ConvertStrSliceToIntSlice(slc []string) []int {
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

func GetSeaBanks(c *gin.Context, trip *models.Trip, ids []int) {
	var sbs []models.SeaBank
	result := db.DB.Find(&sbs, ids)
	if result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Println(err)
		}
	}
	trip.SeaBanks = append(trip.SeaBanks, sbs...)
}

func GetFishTypes(c *gin.Context, trip *models.Trip, ids []int) {
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
	emps := ConvertStrSliceToIntSlice(empStrIDs)
	GetEmployeesForTrip(c, &trip, emps)

	sbsStr := c.PostFormArray("seabanks")
	log.Println(sbsStr)
	sbs := ConvertStrSliceToIntSlice(sbsStr)
	GetSeaBanks(c, &trip, sbs)

	ftsStr := c.PostFormArray("fish")
	log.Println(ftsStr)
	fts := ConvertStrSliceToIntSlice(ftsStr)
	GetFishTypes(c, &trip, fts)

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

func EndTripForm(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	tripId := c.Param("id")
	var trip models.Trip

	if result := db.DB.Preload("FishTypes").Find(&trip, tripId); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "endTrip", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}

	c.HTML(http.StatusOK, "endTrip", gin.H{
		"trip": trip,
		"user": user,
	})
}

func EndTrip(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	strID := c.Param("id")
	tripID, _ := strconv.Atoi(strID)
	var ftIDs []int
	if result := db.DB.Model(&models.FishType{}).Select("id").Find(&ftIDs); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "endTrip", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
	}
	for _, ftID := range ftIDs {
		catch, err := strconv.Atoi(c.PostForm(strconv.Itoa(ftID)))
		if err != nil {
			log.Println(err)
		}
		log.Println("Catch", catch)
		if result := db.DB.Model(&models.FishTypeTrip{}).Where("trip_id = ? and fish_type_id = ?", tripID, ftID).Updates(&models.FishTypeTrip{Catch: catch}); result.Error != nil {
			utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
			c.HTML(http.StatusInternalServerError, "endTrip", gin.H{
				"user":   user,
				"alerts": utils.Flashes(c),
			})
			return
		}
	}
	if result := db.DB.Model(&models.Trip{}).Where("id = ?", tripID).Updates(&models.Trip{ArrivalDate: time.Now()}); result.Error != nil {
		utils.FlashMessage(c, "Возникла ошибка при запросе к базе данных")
		c.HTML(http.StatusInternalServerError, "endTrip", gin.H{
			"user":   user,
			"alerts": utils.Flashes(c),
		})
		return
	}
	dest_url := url.URL{Path: fmt.Sprintf("/trips")}
	c.Redirect(http.StatusSeeOther, dest_url.String())
}
