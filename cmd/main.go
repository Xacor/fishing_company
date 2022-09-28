package main

import (
	"fishing_company/pkg/boats"
	"fishing_company/pkg/config"
	"fishing_company/pkg/db"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err.Error())
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
	router.Use(gin.LoggerWithFormatter(config.CustomLogFormatter))
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("ui/html/**/*")

	handler := db.Init(conf.DBUrl)

	boats.RegisterRoutes(router, handler)
	router.Run(conf.Port)
}
