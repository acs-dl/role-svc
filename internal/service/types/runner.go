package types

import (
	"context"

	"github.com/acs-dl/role-svc/internal/config"
)

type Runner = func(context context.Context, config config.Config)
