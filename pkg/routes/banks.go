package routes

import (
	"fishing_company/pkg/controllers"
	"fishing_company/pkg/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func bankRoutes(superRoute *gin.RouterGroup, authEnforcer *casbin.Enforcer, isTesting bool) {

	bankRouter := superRoute.Group("/banks")
	if !isTesting {
		bankRouter.Use(middleware.AuthRequired, middleware.Authorization(authEnforcer))
	}

	bankRouter.GET("", controllers.GetBanks)
	bankRouter.GET("/:id", controllers.GetBank)
	bankRouter.GET("/create", controllers.BankForm)
	bankRouter.POST("/create", controllers.CreateBank)
	bankRouter.POST("/:id/delete", controllers.DeleteBank)
}
