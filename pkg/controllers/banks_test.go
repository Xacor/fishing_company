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

func TestGetBanks(t *testing.T) {
	router := setupRouter()
	var tests = []struct {
		name     string
		method   string
		path     string
		wantCode int
	}{
		{"GET should work", http.MethodGet, "/banks/", http.StatusOK},
		{"Other methods should not work", http.MethodPost, "/banks/", http.StatusNotFound},
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

func TestGetBank(t *testing.T) {
	router := setupRouter()
	var tests = []struct {
		name     string
		method   string
		path     string
		wantCode int
	}{
		{"GET should work", http.MethodGet, "/banks/1", http.StatusOK},
		{"Other methods should not work", http.MethodPost, "/banks/1", http.StatusNotFound},
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

func TestBankForm(t *testing.T) {
	router := setupRouter()
	var tests = []struct {
		name     string
		method   string
		path     string
		wantCode int
	}{
		{"GET Bank Form", http.MethodGet, "/banks/create", http.StatusOK},
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

func TestCreateBank(t *testing.T) {
	router := setupRouter()
	type bank struct {
		Lat  string
		Lng  string
		Name string
	}

	var tests = []struct {
		name     string
		method   string
		path     string
		data     bank
		wantCode int
	}{
		{"POST Bank Form", http.MethodPost, "/banks/create",
			bank{
				Name: "TestSeaBank",
				Lat:  "49.25",
				Lng:  "55.00",
			},
			http.StatusSeeOther,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{
				"name": {tt.data.Name},
				"lat":  {tt.data.Lat},
				"lng":  {tt.data.Lng},
			}

			req, _ := http.NewRequest(tt.method, tt.path, bytes.NewBuffer([]byte(form.Encode())))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantCode, w.Code)
			db.DB.Where("name = ?", tt.data.Name).Delete(&models.SeaBank{})
		})
	}
}

func TestDeleteBank(t *testing.T) {
	router := setupRouter()

	type bank struct {
		Lat  float64
		Lng  float64
		Name string
	}
	var tests = []struct {
		name     string
		method   string
		path     string
		bank     bank
		wantCode int
	}{
		{"Delete Bank", http.MethodPost, "/banks/999/delete",
			bank{
				Name: "TestBank",
				Lat:  10.10,
				Lng:  80.00,
			},
			http.StatusSeeOther,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := models.SeaBank{ID: 999, Name: tt.bank.Name, Lat: tt.bank.Lat, Lng: tt.bank.Lng}
			if result := db.DB.Create(&bank); result.Error != nil {
				t.Log(result.Error)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, tt.path, nil)

			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantCode, w.Code)

		})
	}
}
