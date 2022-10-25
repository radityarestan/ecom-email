package pkg

import (
	"github.com/nsqio/go-nsq"
	"github.com/radityarestan/ecom-email/internal/config"
	"github.com/radityarestan/ecom-email/internal/service"
)

type NSQConsumer struct {
	Consumer  *nsq.Consumer
	Consumer2 *nsq.Consumer
	Env       config.NSQConfig
}

func (nc *NSQConsumer) Start(emailSender, passSender string) error {
	nc.Consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		go service.SendEmailVerification(emailSender, passSender, string(message.Body))
		message.Finish()
		return nil
	}))

	nc.Consumer2.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		// TODO: implement this
		message.Finish()
		return nil
	}))

	if err := nc.Consumer.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return err
	}

	if err := nc.Consumer2.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return err
	}

	return nil
}

func (nc *NSQConsumer) Stop() {
	nc.Consumer.Stop()
}
