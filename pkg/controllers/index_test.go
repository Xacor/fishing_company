package controllers_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Xacor/fishing_company/pkg/config"
	"github.com/Xacor/fishing_company/pkg/db"
	"github.com/Xacor/fishing_company/pkg/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	conf, err := config.LoadConfig("../../envs")
	if err != nil {
		log.Fatalln(err.Error())
	}

	router := gin.New()

	store := cookie.NewStore([]byte(conf.Secret))
	router.Use(sessions.Sessions("session", store))
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	routes.RegisterRoutes(&router.RouterGroup, nil, true)
	router.LoadHTMLGlob("../../ui/html/*/*.html")
	router.Static("/static", "../../ui/static")

	db.Init(conf.TestDBUrl)

	return router

}

func TestIndex(t *testing.T) {
	router := setupRouter()

	var tests = []struct {
		name     string
		method   string
		path     string
		wantCode int
	}{
		{"GET should work", "GET", "/", http.StatusOK},
		{"POST should not work", "POST", "/", http.StatusNotFound},
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
