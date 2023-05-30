package main

import (
	"os"

	"github.com/Xacor/fishing_company/pkg/config"
	"github.com/Xacor/fishing_company/pkg/db"
	"github.com/Xacor/fishing_company/pkg/logger"
	"github.com/Xacor/fishing_company/pkg/middleware"
	"github.com/Xacor/fishing_company/pkg/routes"
	log "github.com/sirupsen/logrus"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "timestamp",
			log.FieldKeyLevel: "level",
			log.FieldKeyMsg:   "message",
			log.FieldKeyFunc:  "caller",
		},
	})

	conf, err := config.LoadConfig("./envs")
	if err != nil {
		log.Fatalln(err)
	}

	// REPLACE HARDCODED URL
	log.AddHook(logger.NewHook(conf.LoggingURL))

	authEnforcer, err := casbin.NewEnforcer("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatalln(err)
	}

	router := gin.New()
	store := cookie.NewStore([]byte(conf.Secret))
	router.Use(sessions.Sessions("session", store))
	router.Use(middleware.Logger(conf.LoggingURL), middleware.Prometheus)
	router.Use(gin.Recovery())
	routes.RegisterRoutes(&router.RouterGroup, authEnforcer, false)
	router.LoadHTMLGlob("ui/html/*/*.html")
	router.Static("/static", "./ui/static")

	db.Init(conf.DBUrl)

	log.Info("started serving")
	if err := router.Run(conf.Port); err != nil {
		log.Fatalln(err)
	}

}
