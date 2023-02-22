package routes

import (
	"github.com/Xacor/fishing_company/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func prometheusRoutes(superRoute *gin.RouterGroup) {
	prometheus.Register(middleware.TotalRequests)
	prometheus.Register(middleware.ResponseStatus)

	prometheusRouter := superRoute.Group("/")
	prometheusRouter.GET("/metrics", gin.WrapH(promhttp.Handler()))

}
