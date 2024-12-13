package suite

import (
	"context"
	"firstcode-auth/internal/config"
	"testing"

	ssov1 "github.com/MirzaDgtu/firstcode_protos/gen/go/sso"
)

type Suite struct {
	*testing.T
	Cfg        *config.Config
	AuthClient ssov1.AuthClient
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

}
