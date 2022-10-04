package configuration_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zuluapp/go-libs/pkg/libs/core/context"
	"github.com/zuluapp/go-libs/pkg/libs/core/entities"
	"github.com/zuluapp/go-libs/pkg/libs/metrics/prometheus"
	"github.com/zuluapp/prometheus-go/internal/infraestructure/configuration"
)

func Test_GetConfiguration(t *testing.T) {
	ass := assert.New(t)

	context.InitCustomTestContext("1234567890", entities.EnvProduction, entities.RoleWorker, "0.0.1", prometheus.ManagerName())

	conf := configuration.GetConfiguration()
	ass.NotNil(conf)
	ass.Implements((*configuration.Configuration)(nil), conf)
}

func Test_Configuration_Dev(t *testing.T) {
	ass := assert.New(t)

	context.InitCustomTestContext("1234567890", entities.EnvDevelop, entities.RoleWorker, "0.0.1", prometheus.ManagerName())

	conf := configuration.GetConfiguration()

	ass.NotNil(conf)
	ass.Implements((*configuration.Configuration)(nil), conf)
}
