package routes

import (
	"fishing_company/pkg/controllers"
	"fishing_company/pkg/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func fishRoutes(superRoute *gin.RouterGroup, authEnforcer *casbin.Enforcer, isTesting bool) {

	fishRouter := superRoute.Group("/fishes")

	if !isTesting {
		fishRouter.Use(middleware.AuthRequired, middleware.Authorization(authEnforcer))
	}

	fishRouter.GET("/", controllers.GetFishes)
	fishRouter.GET("/create", controllers.FishForm)
	fishRouter.POST("/create", controllers.CreateFish)
	fishRouter.POST("/:id/delete", controllers.DeleteFish)
}
