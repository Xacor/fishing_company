package controllers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Xacor/fishing_company/pkg/db"
	"github.com/Xacor/fishing_company/pkg/models"

	"github.com/stretchr/testify/assert"
)

func TestGetFishes(t *testing.T) {
	router := setupRouter()
	var tests = []struct {
		name     string
		method   string
		path     string
		wantCode int
	}{
		{"GET should work", http.MethodGet, "/fishes/", http.StatusOK},
		{"Other methods should not work", http.MethodPost, "/fishes/", http.StatusNotFound},
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

func TestFishForm(t *testing.T) {
	router := setupRouter()
	var tests = []struct {
		name     string
		method   string
		path     string
		wantCode int
	}{
		{"GET", http.MethodGet, "/fishes/create", http.StatusOK},
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

func TestCreateFish(t *testing.T) {
	router := setupRouter()

	var tests = []struct {
		name     string
		method   string
		path     string
		fishName string
		wantCode int
	}{
		{"Good POST", http.MethodPost, "/fishes/create", "TestName", http.StatusSeeOther},
		//{"Bad POST, fish type exist", http.MethodPost, "/fishes/create", "TestName", http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{"name": {tt.fishName}}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, tt.path, bytes.NewBuffer([]byte(form.Encode())))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantCode, w.Code)

		})
	}
	db.DB.Where("name = ?", "TestName").Delete(&models.FishType{})
}

func TestDeleteFish(t *testing.T) {
	router := setupRouter()

	fish := models.FishType{ID: 999, Name: "TestName"}
	if result := db.DB.Create(&fish); result.Error != nil {
		t.Log(result.Error)
	}

	var tests = []struct {
		name     string
		method   string
		path     string
		fishID   string
		wantCode int
	}{
		{"Delete fish", http.MethodPost, "/fishes/999/delete", "999", http.StatusSeeOther},
		//{"Bad DELETE, fish type not exist", http.MethodPost, "/fishes/1000/delete", "1000", http.StatusInternalServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, tt.path, nil)

			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantCode, w.Code)

		})
	}
}
