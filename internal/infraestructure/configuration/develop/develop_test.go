package develop

import (
	"os"
	"testing"

	"github.com/zuluapp/go-libs/pkg/constants"
	"github.com/zuluapp/go-libs/pkg/libs/core/context"
)

var (
	Application      = "TODO"
	ApplicationID    = "1234567890"
	ApplicationScope = "test-read"
	Environment      = "TEST"
)

func TestMain(m *testing.M) {
	os.Setenv(constants.Application, Application)
	os.Setenv(constants.ApplicationID, ApplicationID)
	os.Setenv(constants.ApplicationScope, ApplicationScope)
	os.Setenv(constants.Environment, Environment)

	context.New()

	os.Exit(m.Run())
}
