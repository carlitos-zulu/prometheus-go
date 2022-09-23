package main

import (
	"github.com/gin-gonic/gin"

	"github.com/zuluapp/go-libs/pkg/libs/core"
	"github.com/zuluapp/go-libs/pkg/libs/core/application"
	"github.com/zuluapp/go-libs/pkg/libs/core/entities"

	"github.com/zuluapp/prometheus-go/internal/infraestructure/dependencies"
	"github.com/zuluapp/prometheus-go/internal/infraestructure/roles/read"
)

type contextBuilder struct{}

const basePath = "/zulu-metrics"

var builder core.BuildContext = contextBuilder{}

func main() {
	application.Start(builder)
}

func (contextBuilder) GetRoutingGroup() entities.RoutingGroup {
	container := dependencies.StartDependencies()

	return entities.RoutingGroup{
		entities.RoleRead: func() func(*gin.RouterGroup) { return read.New(container).RegisterRoutes(basePath) },
	}
}

func (contextBuilder) GetGroupPrefix() string {
	return ""
}
