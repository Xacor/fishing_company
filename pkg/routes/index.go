package routes

import (
	"github.com/Xacor/fishing_company/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func indexRoutes(superRoute *gin.RouterGroup) {

	boatRouter := superRoute.Group("/")

	boatRouter.GET("/", controllers.Index)

}
