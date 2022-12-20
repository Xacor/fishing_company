package controllers_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/Xacor/fishing_company/pkg/db"
	"github.com/Xacor/fishing_company/pkg/models"

	"github.com/stretchr/testify/assert"
)

func TestGetEmployees(t *testing.T) {
	router := setupRouter()
	var tests = []struct {
		name     string
		method   string
		path     string
		wantCode int
	}{
		{"GET should work", "GET", "/employees/", http.StatusOK},
		{"POST should not work", "POST", "/employees/", http.StatusNotFound},
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

func TestGetEmployee(t *testing.T) {
	router := setupRouter()

	var tests = []struct {
		name     string
		method   string
		path     string
		wantCode int
	}{
		{"GET to existing employee", "GET", "/employees/6", http.StatusOK},
		{"GET to  not existing employee", "GET", "/employees/999", http.StatusInternalServerError},
		{"POST should not work", "POST", "/employees/", http.StatusNotFound},
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

func TestEmployeeForm(t *testing.T) {
	router := setupRouter()

	var tests = []struct {
		name     string
		method   string
		path     string
		wantCode int
	}{
		{"GET to createEmployeeForm", "GET", "/employees/create", http.StatusOK},
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

func TestCreateEmployee(t *testing.T) {
	router := setupRouter()

	var tests = []struct {
		name        string
		method      string
		path        string
		bLastname   string
		bFirstname  string
		bMiddlename string
		bAddress    string
		bBirth_date string
		bPositionID string
		wantCode    int
	}{
		{"POST ",
			"POST",
			"/employees/create",
			"Sidorov",
			"Sidr",
			"",
			"Moscow",
			"2010-10-10",
			"1",
			http.StatusSeeOther},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{"lastname": {tt.bLastname}, "firstname": {tt.bFirstname}, "middlename": {tt.bMiddlename}, "address": {tt.bAddress}, "birth_date": {tt.bBirth_date}, "position": {tt.bPositionID}}
			req, _ := http.NewRequest(tt.method, tt.path, bytes.NewBuffer([]byte(form.Encode())))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantCode, w.Code)
		})
	}
}

func TestDeleteEmployee(t *testing.T) {
	router := setupRouter()

	rand.Seed(time.Now().UnixNano())
	min := 30000
	max := 30100
	eID := rand.Intn(max-min+1) + min
	emp := models.Employee{
		ID:         eID,
		Lastname:   fmt.Sprintf("TestEmployee_%v", eID),
		Firstname:  fmt.Sprintf("TestEmployee_%v", eID),
		Middlename: fmt.Sprintf("TestEmployee_%v", eID),
		PositionID: 1,
		Address:    "MosPolytech",
		Birth_date: time.Now(),
	}
	if result := db.DB.Create(&emp); result.Error != nil {
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
			fmt.Sprintf("/employees/%v/delete", eID),
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

func TestUpdateEmployeeForm(t *testing.T) {
	router := setupRouter()

	var tests = []struct {
		name     string
		method   string
		path     string
		wantCode int
	}{
		{"GET to updateEmployeeForm", "GET", "/employees/6/update", http.StatusOK},
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

func TestUpdateEmployee(t *testing.T) {
	router := setupRouter()

	rand.Seed(time.Now().UnixNano())
	min := 30200
	max := 30300
	eID := rand.Intn(max-min+1) + min
	emp := models.Employee{
		ID:         eID,
		Lastname:   fmt.Sprintf("TestEmployee_%v", eID),
		Firstname:  fmt.Sprintf("TestEmployee_%v", eID),
		Middlename: fmt.Sprintf("TestEmployee_%v", eID),
		PositionID: 1,
		Address:    "MosPolytech",
		Birth_date: time.Now(),
	}
	if result := db.DB.Create(&emp); result.Error != nil {
		t.Log(result.Error)
	}

	var tests = []struct {
		name        string
		method      string
		path        string
		bLastname   string
		bFirstname  string
		bMiddlename string
		bAddress    string
		bBirth_date string
		bPositionID string
		wantCode    int
	}{
		{"POST update employee",
			"POST",
			fmt.Sprintf("/employees/%v/update", eID),
			fmt.Sprintf("TestEmployee_UPD%v", eID),
			fmt.Sprintf("TestEmployee_UPD%v", eID),
			fmt.Sprintf("TestEmployee_UPD%v", eID),
			"Moscow",
			"2010-10-10",
			"1",
			http.StatusSeeOther},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{"lastname": {tt.bLastname}, "firstname": {tt.bFirstname}, "middlename": {tt.bMiddlename}, "address": {tt.bAddress}, "birth_date": {tt.bBirth_date}, "position": {tt.bPositionID}}
			req, _ := http.NewRequest(tt.method, tt.path, bytes.NewBuffer([]byte(form.Encode())))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantCode, w.Code)
		})
	}
}
