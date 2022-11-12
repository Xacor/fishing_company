package main

import (
	"fishing_company/pkg/config"
	"fishing_company/pkg/db"
	"fishing_company/pkg/routes"
	"io"
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

//random comment

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

	router := gin.New()

	//replace with normal auth key
	store := memstore.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))

	router.Use(gin.LoggerWithFormatter(config.CustomLogFormatter))
	router.Use(gin.Recovery())
	routes.RegisterRoutes(&router.RouterGroup)
	router.LoadHTMLGlob("ui/html/*/*.html")
	//router.LoadHTMLGlob("ui/html/**/*")

	db.Init(conf.DBUrl)

	if err := router.Run(conf.Port); err != nil {
		log.Fatalln(err)
	}
}
