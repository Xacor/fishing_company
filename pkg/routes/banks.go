package routes

import (
	"fishing_company/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func bankRoutes(superRoute *gin.RouterGroup) {

	bankRouter := superRoute.Group("/banks")

	bankRouter.GET("/", controllers.GetBanks)
	bankRouter.GET("/:id", controllers.GetBank)
	bankRouter.GET("/create", controllers.BankForm)
	bankRouter.POST("/create", controllers.CreateBank)
	bankRouter.GET("/:id/delete", controllers.DeleteBank)
}
