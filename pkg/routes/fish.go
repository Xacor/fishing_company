package routes

import (
	"fishing_company/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func fishRoutes(superRoute *gin.RouterGroup) {

	fishRouter := superRoute.Group("/fishes")
	fishRouter.GET("/", controllers.GetFishes)
	fishRouter.GET("/create", controllers.FishForm)
	fishRouter.POST("/create", controllers.CreateFish)
	fishRouter.POST("/:id/delete", controllers.DeleteFish)
}
