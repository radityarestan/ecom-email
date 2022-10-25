package shared

import (
	"github.com/radityarestan/ecom-email/internal/pkg"
	"github.com/radityarestan/ecom-email/internal/shared/config"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type Deps struct {
	dig.In
	Config      *config.Config
	Logger      *log.Logger
	NSQConsumer *pkg.NSQConsumer
}

func (d *Deps) Close() {
	d.Logger.Info("Closing NSQ Consumer")
	d.NSQConsumer.Stop()
}
