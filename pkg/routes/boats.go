package routes

import (
	"fishing_company/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func boatRoutes(superRoute *gin.RouterGroup) {

	boatRouter := superRoute.Group("/boats")

	boatRouter.GET("/", controllers.GetBoats)
	boatRouter.GET("/create", controllers.BoatForm)
	boatRouter.POST("/create", controllers.CreateBoat)
	boatRouter.GET("/:id", controllers.GetBoat)
	boatRouter.GET("/:id/update", controllers.UpdateBoatForm)
	boatRouter.POST("/:id/update", controllers.UpdateBoat)
	boatRouter.POST("/:id/delete", controllers.DeleteBoat)
}
