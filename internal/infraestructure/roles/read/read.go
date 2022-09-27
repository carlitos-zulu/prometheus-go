package read

import (
	"github.com/gin-gonic/gin"
	"github.com/zuluapp/go-libs/pkg/utils"
	"github.com/zuluapp/prometheus-go/internal/infraestructure/dependencies"
	prom "github.com/zuluapp/prometheus-go/internal/infraestructure/prometheus"
)

type readContainer struct {
	container *dependencies.Container
}

func New(container *dependencies.Container) *readContainer {
	return &readContainer{
		container: container,
	}
}

func (container readContainer) RegisterRoutes(basePath string) func(*gin.RouterGroup) {
	return func(g *gin.RouterGroup) {
		prom.Setup(g)

		v1Group := g.Group(basePath + "/v1")
		roleGroup := v1Group.Group("/test")

		metrics := container.container.Metrics()

		roleGroup.GET("/", func(ctx *gin.Context) {
			logDuration := metrics.ObserveDuration("zulu_my_custom_duration", utils.Map{
				"carlitos": "otro",
			})
			defer logDuration()

			metrics.Increment("zulu_my_custom_metric", utils.Map{
				"carlitos": "carlitos-zulu",
			})

			metrics.IncrementGauge("zulu_my_custom_gauge", utils.Map{
				"carlitos": "gauge",
			})

			response := utils.Map{
				"message": "funciona!",
			}
			ctx.JSON(200, response)
		})
	}
}
