package main

import (
	"fishing_company/pkg/common/boats"
	"fishing_company/pkg/common/config"
	"fishing_company/pkg/common/db"

	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err.Error())
	}

	router := gin.Default()
	router.LoadHTMLGlob("ui/html/**/*")

	handler := db.Init(conf.DBUrl)

	boats.RegisterRoutes(router, handler)

	router.Run(conf.Port)
}
