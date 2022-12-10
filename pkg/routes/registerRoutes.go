package routes

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(superRoute *gin.RouterGroup, enforcer *casbin.Enforcer) {
	boatRoutes(superRoute, enforcer)
	authRoutes(superRoute, enforcer)
	indexRoutes(superRoute)
	bankRoutes(superRoute, enforcer)
	fishRoutes(superRoute, enforcer)
	employeeRoutes(superRoute, enforcer)
}
