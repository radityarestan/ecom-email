package di

import (
	"github.com/radityarestan/ecom-email/internal/config"
	"github.com/radityarestan/ecom-email/internal/pkg"
)

type Deps struct {
	Config      *config.Config
	NSQConsumer *pkg.NSQConsumer
}

func InitDeps() (*Deps, error) {
	config, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	nsqConsumer, err := NewNSQConsumer(config)
	if err != nil {
		return nil, err
	}

	return &Deps{
		Config:      config,
		NSQConsumer: nsqConsumer,
	}, nil
}
