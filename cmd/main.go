package main

import (
	"fishing_company/pkg/auth"
	"fishing_company/pkg/boats"
	"fishing_company/pkg/config"
	"fishing_company/pkg/db"
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
	router.LoadHTMLGlob("ui/html/**/*")

	db.Init(conf.DBUrl)

	boats.RegisterRoutes(router)
	auth.RegisterRoutes(router)
	if err := router.Run(conf.Port); err != nil {
		log.Fatalln(err)
	}
}
