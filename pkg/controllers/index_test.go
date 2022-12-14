package controllers_test

import (
	"fishing_company/pkg/config"
	"fishing_company/pkg/db"
	"fishing_company/pkg/routes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

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

	switch conf.LogO {
	case "file":
		gin.DisableConsoleColor()
		f, _ := os.Create(conf.LogFile)
		gin.DefaultWriter = io.MultiWriter(f)
	case "all":
		f, _ := os.Create(conf.LogFile)
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}

	router := gin.New()

	store := cookie.NewStore([]byte(conf.Secret))
	router.Use(sessions.Sessions("session", store))
	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: ioutil.Discard}))

	routes.RegisterRoutes(&router.RouterGroup)
	router.LoadHTMLGlob("../../ui/html/*/*.html")
	router.Static("/static", "../../ui/static")

	db.Init(conf.DBUrl)

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
