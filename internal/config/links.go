package config

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type LinksCfg struct {
	Orchestrator string `fig:"orchestrator,required"`
}

func (c *config) Links() *LinksCfg {
	return c.links.Do(func() interface{} {
		cfg := LinksCfg{}

		if err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "links")).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to get links from config"))
		}

		err := cfg.validate()
		if err != nil {
			panic(errors.Wrap(err, "failed to validate links config"))
		}

		return &cfg
	}).(*LinksCfg)
}

func (lc *LinksCfg) validate() error {
	return validation.Errors{
		"orchestrator_link": validation.Validate(&lc.Orchestrator, validation.Required),
	}.Filter()
}
