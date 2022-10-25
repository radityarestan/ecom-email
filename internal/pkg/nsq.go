package pkg

import (
	"github.com/nsqio/go-nsq"
	"github.com/radityarestan/ecom-email/internal/shared/config"
)

type NSQConsumer struct {
	Consumer  *nsq.Consumer
	Consumer2 *nsq.Consumer
	Env       config.NSQConfig
}

func (nc *NSQConsumer) Stop() {
	nc.Consumer.Stop()
}
