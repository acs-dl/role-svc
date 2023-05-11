package service

import (
	"context"
	"sync"

	"gitlab.com/distributed_lab/acs/role-svc/internal/config"
	"gitlab.com/distributed_lab/acs/role-svc/internal/registrator"
	"gitlab.com/distributed_lab/acs/role-svc/internal/service/api"
	"gitlab.com/distributed_lab/acs/role-svc/internal/service/types"
)

var availableServices = map[string]types.Runner{
	"api":       api.Run,
	"registrar": registrator.Run,
}

func Run(cfg config.Config) {
	logger := cfg.Log().WithField("service", "main")
	ctx := context.Background()
	wg := new(sync.WaitGroup)

	logger.Info("Starting all available services...")

	// create new tg sessions better from this point
	//tgClient := tg.NewTg(cfg.Telegram(), cfg.Log())

	for serviceName, service := range availableServices {
		wg.Add(1)

		go func(name string, runner types.Runner) {
			defer wg.Done()

			runner(ctx, cfg)

		}(serviceName, service)

		logger.WithField("service", serviceName).Info("Service started")
	}

	wg.Wait()

}
