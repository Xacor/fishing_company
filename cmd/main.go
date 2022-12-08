package main

import (
	"fishing_company/pkg/config"
	"fishing_company/pkg/db"
	"fishing_company/pkg/routes"
	"io"
	"log"
	"os"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.LoadConfig()
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

	authEnforcer, err := casbin.NewEnforcer("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatalln(err)
	}

	router := gin.New()

	store := cookie.NewStore([]byte(conf.Secret))
	router.Use(sessions.Sessions("session", store))

	router.Use(gin.LoggerWithFormatter(config.CustomLogFormatter))
	router.Use(gin.Recovery())
	routes.RegisterRoutes(&router.RouterGroup, authEnforcer)
	router.LoadHTMLGlob("ui/html/*/*.html")
	router.Static("/static", "./ui/static")
	//router.LoadHTMLGlob("ui/html/**/*")

	db.Init(conf.DBUrl)

	if err := router.Run(conf.Port); err != nil {
		log.Fatalln(err)
	}
}
