package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "status"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}

func RegisterMetrics(router *gin.Engine) {
	router.GET("api/v1/metrics", gin.WrapH(promhttp.Handler()))
}
