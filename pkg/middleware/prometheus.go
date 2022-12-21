package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var TotalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path"},
)

var ResponseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status",
		Help: "Status of HTTP response",
	},
	[]string{"status"},
)

func Prometheus(c *gin.Context) {
	path := c.Request.URL.Path
	TotalRequests.WithLabelValues(path).Inc()

	code := c.Writer.Status()
	ResponseStatus.WithLabelValues(strconv.Itoa(code)).Inc()

	c.Next()
}
