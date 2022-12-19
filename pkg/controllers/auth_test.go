package controllers_test

import (
	"bytes"
	"fishing_company/pkg/db"
	"fishing_company/pkg/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginForm(t *testing.T) {
	router := setupRouter()

	var tests = []struct {
		name     string
		method   string
		path     string
		wantCode int
	}{
		{"GET should work", http.MethodGet, "/auth/login", http.StatusOK},
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

func TestRegisterForm(t *testing.T) {
	router := setupRouter()

	var tests = []struct {
		name     string
		method   string
		path     string
		wantCode int
	}{
		{"GET should work", http.MethodGet, "/auth/register", http.StatusOK},
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

func TestRegister(t *testing.T) {
	router := setupRouter()
	type creds struct {
		Username     string
		Password     string
		ConfPassword string
		Role         string
	}
	var tests = []struct {
		name     string
		method   string
		path     string
		creds    creds
		wantCode int
	}{
		{"Register admin", http.MethodPost, "/auth/register",
			creds{
				Username:     "test",
				Password:     "qwerty",
				ConfPassword: "qwerty",
				Role:         "1",
			},
			http.StatusSeeOther,
		},
		{"Register user", http.MethodPost, "/auth/register",
			creds{
				Username:     "test2",
				Password:     "qwerty",
				ConfPassword: "qwerty",
				Role:         "2",
			},
			http.StatusSeeOther,
		},
		{"Wrong passwords", http.MethodPost, "/auth/register",
			creds{
				Username:     "test3",
				Password:     "qwer",
				ConfPassword: "qwerty",
				Role:         "1",
			},
			http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{
				"username":      {tt.creds.Username},
				"password":      {tt.creds.Password},
				"conf-password": {tt.creds.ConfPassword},
				"role":          {tt.creds.Role},
			}
			req, _ := http.NewRequest(tt.method, tt.path, bytes.NewBuffer([]byte(form.Encode())))
			w := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantCode, w.Code)
			db.DB.Where("name = ?", tt.creds.Username).Delete(&models.User{})
		})
	}
}
