package dependencies

import (
	"github.com/zuluapp/go-libs/pkg/libs/metrics"
	"github.com/zuluapp/go-libs/pkg/libs/metrics/entities"

	"github.com/zuluapp/prometheus-go/internal/infraestructure/configuration"
)

type Container struct {
	metrics entities.Metric
}

func StartDependencies() *Container {
	/* config := */ configuration.GetConfiguration()

	metrics := metrics.Get()

	return &Container{
		metrics: metrics,
	}
}

func (c Container) Metrics() entities.Metric {
	return c.metrics
}
