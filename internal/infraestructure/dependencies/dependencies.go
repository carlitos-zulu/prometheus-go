package dependencies

import (
	"github.com/zuluapp/prometheus-go/internal/infraestructure/configuration"
	"github.com/zuluapp/prometheus-go/internal/infraestructure/prometheus"
)

type Container struct {
	metrics prometheus.Metrics
}

func StartDependencies() *Container {
	/* config := */ configuration.GetConfiguration()

	metrics := prometheus.InitMetrics()

	return &Container{
		metrics: metrics,
	}
}

func (c Container) Metrics() prometheus.Metrics {
	return c.metrics
}
