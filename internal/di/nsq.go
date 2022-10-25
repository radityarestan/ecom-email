package di

import (
	"github.com/nsqio/go-nsq"
	"github.com/radityarestan/ecom-email/internal/config"
	"github.com/radityarestan/ecom-email/internal/pkg"
)

func NewNSQConsumer(conf *config.Config) (nc *pkg.NSQConsumer, err error) {
	nc = &pkg.NSQConsumer{}
	nc.Env = conf.NSQ

	nsqConfig := nsq.NewConfig()
	nc.Consumer, err = nsq.NewConsumer(nc.Env.Topic, nc.Env.Channel, nsqConfig)
	if err != nil {
		return nil, err
	}

	nc.Consumer2, err = nsq.NewConsumer(nc.Env.Topic2, nc.Env.Channel2, nsqConfig)
	if err != nil {
		return nil, err
	}

	return nc, nil
}
