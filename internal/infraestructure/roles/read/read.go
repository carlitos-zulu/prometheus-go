package read

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/zuluapp/go-libs/pkg/utils"
	utilsHttp "github.com/zuluapp/go-libs/pkg/utils/http"

	"github.com/zuluapp/prometheus-go/internal/infraestructure/dependencies"
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
		v1Group := g.Group(basePath + "/v1")
		roleGroup := v1Group.Group("/test")

		metrics := container.container.Metrics()

		roleGroup.GET("/cash-in", func(ctx *gin.Context) {
			rand.Seed(time.Now().UnixNano())

			quantity, _ := utilsHttp.ExtractIDsParam(ctx, "quantity")

			user := utils.Map{
				"user":     "62cc266b805543b5a0115ce3",
				"quantity": strconv.FormatUint(*quantity, 10),
			}

			logDuration := metrics.ObserveDuration("cash_in_duration", user)
			defer logDuration()

			metrics.Increment("cash_in_quantity", user)

			sleepMax := 1200
			sleepMin := 100

			sleep := rand.Intn(sleepMax-sleepMin+1) + sleepMin

			time.Sleep(time.Duration(sleep) * time.Millisecond)

			ctx.JSON(200, utils.Map{
				"message":  "cash-in-ok",
				"sleeped":  sleep,
				"quantity": quantity,
			})
		})

		roleGroup.GET("/cash-out", func(ctx *gin.Context) {
			rand.Seed(time.Now().UnixNano())

			quantity, _ := utilsHttp.ExtractIDsParam(ctx, "quantity")

			user := utils.Map{
				"user":     "62cc266b805543b5a0115ce3",
				"quantity": strconv.FormatUint(*quantity, 10),
			}

			logDuration := metrics.ObserveDuration("cash_out_duration", user)
			defer logDuration()

			metrics.Increment("cash_out_quantity", user)

			sleepMax := 1200
			sleepMin := 100

			sleep := rand.Intn(sleepMax-sleepMin+1) + sleepMin

			time.Sleep(time.Duration(sleep) * time.Millisecond)

			ctx.JSON(200, utils.Map{
				"message":  "cash-out-ok",
				"sleeped":  sleep,
				"quantity": quantity,
			})
		})
	}
}
