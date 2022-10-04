package dependencies

import (
	"github.com/zuluapp/go-libs/pkg/libs/metrics"
	"github.com/zuluapp/go-libs/pkg/libs/metrics/entities"

	"github.com/zuluapp/prometheus-go/internal/infraestructure/configuration"
)

type Container struct{}

func StartDependencies() *Container {
	/* config := */ configuration.GetConfiguration()

	return &Container{}
}

func (Container) Metrics() entities.Metric {
	return metrics.Get()
}
