package di

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/radityarestan/ecom-email/internal/pkg"
	"github.com/radityarestan/ecom-email/internal/shared/config"
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

	nc.Consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Println("Received a message (CONSUMER 1): ", string(message.Body))
		message.Finish()
		return nil
	}))

	nc.Consumer2.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Println("Received a message (CONSUMER 2): ", string(message.Body))
		message.Finish()
		return nil
	}))

	if err := nc.Consumer.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return nil, err
	}

	if err := nc.Consumer2.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return nil, err
	}

	return nc, nil
}
