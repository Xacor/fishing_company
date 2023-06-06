package middleware

import (
	"strconv"
	"time"

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
var ResponseTimeHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "http_server_request_duration_seconds",
	Help:    "Histogram of response time for handler in seconds",
	Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
}, []string{"route", "method", "status_code"})

func Prometheus(c *gin.Context) {
	path := c.Request.URL.Path
	TotalRequests.WithLabelValues(path).Inc()

	code := c.Writer.Status()
	ResponseStatus.WithLabelValues(strconv.Itoa(code)).Inc()
	start := time.Now()
	c.Next()
	duration := time.Since(start)
	ResponseTimeHistogram.WithLabelValues(
		c.Request.RequestURI,
		c.Request.Method,
		strconv.Itoa(c.Writer.Status())).Observe(duration.Seconds())
}
