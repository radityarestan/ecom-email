package di

import (
	"github.com/radityarestan/ecom-email/internal/shared/config"
	"go.uber.org/dig"
)

var (
	Container = dig.New()
)

func init() {
	if err := Container.Provide(config.NewConfig); err != nil {
		panic(err)
	}

	if err := Container.Provide(NewLogger); err != nil {
		panic(err)
	}

	if err := Container.Provide(NewNSQConsumer); err != nil {
		panic(err)
	}
}
