package routes

import (
	"github.com/Xacor/fishing_company/pkg/controllers"
	"github.com/Xacor/fishing_company/pkg/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func tripRoutes(superRoute *gin.RouterGroup, authEnforcer *casbin.Enforcer, isTesting bool) {

	tripRouter := superRoute.Group("/trips")

	if !isTesting {
		tripRouter.Use(middleware.AuthRequired, middleware.Authorization(authEnforcer))
	}
	tripRouter.GET("/", controllers.GetTrips)
	tripRouter.GET("/:id", controllers.GetTrip)
	tripRouter.GET("/create", controllers.TripForm)
	tripRouter.POST("/create", controllers.CreateTrip)
	tripRouter.GET("/:id/end", controllers.EndTripForm)
	tripRouter.POST("/:id/end", controllers.EndTrip)
}
