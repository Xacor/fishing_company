package routes

import (
	"fishing_company/pkg/controllers"
	"fishing_company/pkg/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func fishRoutes(superRoute *gin.RouterGroup, e *casbin.Enforcer) {

	fishRouter := superRoute.Group("/fishes")
	fishRouter.Use(middleware.AuthRequired)

	fishRouter.GET("/", controllers.GetFishes)
	fishRouter.GET("/create", controllers.FishForm)
	fishRouter.POST("/create", controllers.CreateFish)
	fishRouter.DELETE("/:id/delete", controllers.DeleteFish)
}
