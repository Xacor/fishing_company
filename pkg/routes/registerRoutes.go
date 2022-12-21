package routes

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(superRoute *gin.RouterGroup, authEnforcer *casbin.Enforcer, isTesting bool) {
	indexRoutes(superRoute)
	authRoutes(superRoute)
	prometheusRoutes(superRoute)
	boatRoutes(superRoute, authEnforcer, isTesting)
	bankRoutes(superRoute, authEnforcer, isTesting)
	fishRoutes(superRoute, authEnforcer, isTesting)
	employeeRoutes(superRoute, authEnforcer, isTesting)
	tripRoutes(superRoute, authEnforcer, isTesting)
}
