package registrator

import (
	"context"

	"github.com/acs-dl/role-svc/internal/config"
)

func Run(ctx context.Context, cfg config.Config) {
	NewRegistrar(cfg).Run(ctx)
}
