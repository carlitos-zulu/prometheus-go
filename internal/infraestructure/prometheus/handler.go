package prometheus

import (
	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Setup(g *gin.RouterGroup) {
	g.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
