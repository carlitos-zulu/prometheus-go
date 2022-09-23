package configuration

import (
	"github.com/zuluapp/go-libs/pkg/libs/core/context"
	"github.com/zuluapp/go-libs/pkg/libs/core/entities"

	"github.com/zuluapp/prometheus-go/internal/infraestructure/configuration/develop"
	"github.com/zuluapp/prometheus-go/internal/infraestructure/configuration/production"
)

// Configuration interface for configuration service.
type Configuration interface {
}

// GetConfiguration returns the configuration service depending on the scope environment variable.
func GetConfiguration() Configuration {
	if context.Get().Environment() == entities.EnvProduction {
		return production.NewConfiguration()
	}

	return develop.NewConfiguration()
}
