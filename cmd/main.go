package main

import (
	"os"

	"github.com/Xacor/fishing_company/pkg/config"
	"github.com/Xacor/fishing_company/pkg/db"
	"github.com/Xacor/fishing_company/pkg/middleware"
	"github.com/Xacor/fishing_company/pkg/routes"
	log "github.com/sirupsen/logrus"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)

	conf, err := config.LoadConfig("./envs")
	if err != nil {
		log.Fatalln(err)
	}

	authEnforcer, err := casbin.NewEnforcer("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatalln(err)
	}
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	store := cookie.NewStore([]byte(conf.Secret))
	router.Use(sessions.Sessions("session", store))
	router.Use(middleware.Logger)
	router.Use(gin.Recovery())
	routes.RegisterRoutes(&router.RouterGroup, authEnforcer, false)
	router.LoadHTMLGlob("ui/html/*/*.html")
	router.Static("/static", "./ui/static")

	db.Init(conf.DBUrl)

	if err := router.Run(conf.Port); err != nil {
		log.Fatalln(err)
	}
}
