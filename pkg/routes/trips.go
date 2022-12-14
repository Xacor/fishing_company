package routes

import (
	"fishing_company/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func tripRoutes(superRoute *gin.RouterGroup) {

	tripRouter := superRoute.Group("/trips")

	tripRouter.GET("/", controllers.GetTrips)
	tripRouter.GET("/:id", controllers.GetTrip)
	tripRouter.GET("/create", controllers.TripForm)
	tripRouter.POST("/create", controllers.CreateTrip)
	tripRouter.GET("/:id/end", controllers.EndTripForm)
	tripRouter.POST("/:id/end", controllers.EndTrip)
}
