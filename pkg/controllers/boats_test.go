package controllers_test

import (
	"bytes"
	"fishing_company/pkg/db"
	"fishing_company/pkg/models"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetBoats(t *testing.T) {
	router := setupRouter()

	var tests = []struct {
		name     string
		method   string
		path     string
		wantCode int
	}{
		{"GET should work", "GET", "/boats/", http.StatusOK},
		{"POST should not work", "POST", "/boats/", http.StatusNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantCode, w.Code)
		})
	}
}

func TestGetBoat(t *testing.T) {
	router := setupRouter()

	var tests = []struct {
		name     string
		method   string
		path     string
		wantCode int
	}{
		{"GET to existing boat should work", "GET", "/boats/14", http.StatusOK},
		{"GET to  not existing boat should not work", "GET", "/boats/1", http.StatusInternalServerError},
		{"POST should not work", "POST", "/boats/", http.StatusNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantCode, w.Code)
		})
	}
}

func TestBoatForm(t *testing.T) {
	router := setupRouter()

	var tests = []struct {
		name     string
		method   string
		path     string
		wantCode int
	}{
		{"GET to createBoatForm", "GET", "/boats/create", http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantCode, w.Code)
		})
	}
}

func TestCreateBoat(t *testing.T) {
	router := setupRouter()

	var tests = []struct {
		name          string
		method        string
		path          string
		bName         string
		bType         string
		bDisplacement string
		build_date    string
		wantCode      int
	}{
		{"POST with duplicate name",
			"POST",
			"/boats/create",
			"ZXC",
			"1",
			"10",
			"2010-10-10",
			http.StatusInternalServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{"name": {tt.bName}, "type": {tt.bType}, "displacement": {tt.bDisplacement}, "build_date": {tt.build_date}}
			req, _ := http.NewRequest(tt.method, tt.path, bytes.NewBuffer([]byte(form.Encode())))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantCode, w.Code)
		})
	}
}

func TestDeleteBoat(t *testing.T) {
	router := setupRouter()

	rand.Seed(time.Now().UnixNano())
	min := 30000
	max := 30100
	bID := rand.Intn(max-min+1) + min
	boat := models.Boat{
		ID:           bID,
		Name:         fmt.Sprintf("TestBoat_%v", bID),
		BtypeID:      1,
		Displacement: 10,
		Build_date:   time.Now(),
	}
	if result := db.DB.Create(&boat); result.Error != nil {
		t.Log(result.Error)
	}

	var tests = []struct {
		name     string
		method   string
		path     string
		wantCode int
	}{
		{"Delete existing fish",
			"POST",
			fmt.Sprintf("/boats/%v/delete", bID),
			http.StatusSeeOther},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantCode, w.Code)
		})
	}
}

func TestUpdateBoatForm(t *testing.T) {
	router := setupRouter()

	var tests = []struct {
		name     string
		method   string
		path     string
		wantCode int
	}{
		{"GET to updateBoatForm", "GET", "/boats/14/update", http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantCode, w.Code)
		})
	}
}

func TestUpdateBoat(t *testing.T) {
	router := setupRouter()

	rand.Seed(time.Now().UnixNano())
	min := 30200
	max := 30300
	bID := rand.Intn(max-min+1) + min
	boat := models.Boat{
		ID:           bID,
		Name:         fmt.Sprintf("TestUpdateBoat_%v", bID),
		BtypeID:      1,
		Displacement: 10,
		Build_date:   time.Now(),
	}
	if result := db.DB.Create(&boat); result.Error != nil {
		t.Log(result.Error)
	}

	var tests = []struct {
		name          string
		method        string
		path          string
		bName         string
		bType         string
		bDisplacement string
		wantCode      int
	}{
		{"Successful boat update",
			"POST",
			fmt.Sprintf("/boats/%v/update", bID),
			fmt.Sprintf("UpdatedBoat_%v", bID),
			"2",
			"10",
			http.StatusSeeOther},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{"name": {tt.bName}, "type": {tt.bType}, "displacement": {tt.bDisplacement}}
			req, _ := http.NewRequest(tt.method, tt.path, bytes.NewBuffer([]byte(form.Encode())))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantCode, w.Code)
		})
	}
}
