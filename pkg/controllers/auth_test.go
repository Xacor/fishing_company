package controllers_test

import (
	"net/http"
	"net/http/httptest"
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

// func TestRegister(t *testing.T) {
// 	router := setupRouter()

// 	var tests = []struct {
// 		name     string
// 		method   string
// 		path     string
// 		username string
// 		password string
// 		wantCode int
// 	}{
// 		{"POST should work", http.MethodPost, "/auth/register", http.StatusSeeOther},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			req, _ := http.NewRequest(tt.method, tt.path, nil)
// 			w := httptest.NewRecorder()
// 			router.ServeHTTP(w, req)
// 			assert.Equal(t, tt.wantCode, w.Code)
// 		})
// 	}
// }
