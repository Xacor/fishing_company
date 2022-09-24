package main

import (
	"fishing_company/pkg/common/boats"
	"fishing_company/pkg/common/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()
	router.LoadHTMLGlob("ui/html/**/*")

	handler := db.Init(dbUrl)

	boats.RegisterRoutes(router, handler)

	// func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"port":  port,
	// 		"dbUrl": dbUrl,
	// 	})
	// }
	router.Run(port)
}
