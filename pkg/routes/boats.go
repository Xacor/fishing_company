package routes

import (
	"fishing_company/pkg/controllers"
	"fishing_company/pkg/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func boatRoutes(superRoute *gin.RouterGroup, e *casbin.Enforcer) {

	boatRouter := superRoute.Group("/boats")
	boatRouter.Use(middleware.AuthRequired)
	boatRouter.Use(middleware.Authorization(e))

	boatRouter.GET("/", controllers.GetBoats)
	boatRouter.GET("/create", controllers.BoatForm)
	boatRouter.POST("/create", controllers.CreateBoat)
	boatRouter.GET("/:id", controllers.GetBoat)
	boatRouter.GET("/:id/update", controllers.UpdateBoatForm)
	boatRouter.POST("/:id/update", controllers.UpdateBoat)
	boatRouter.POST("/:id/delete", controllers.DeleteBoat)
}
