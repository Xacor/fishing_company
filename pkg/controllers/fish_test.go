package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
