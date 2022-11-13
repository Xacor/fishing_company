package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(superRoute *gin.RouterGroup) {
	boatRoutes(superRoute)
	authRoutes(superRoute)
	indexRoutes(superRoute)
	bankRoutes(superRoute)
	fishRoutes(superRoute)
}
