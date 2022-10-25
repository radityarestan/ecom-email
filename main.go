package main

import (
	"github.com/radityarestan/ecom-email/internal/di"
	"github.com/radityarestan/ecom-email/internal/shared"
	"os"
	"os/signal"
)

func main() {
	var container = di.Container

	err := container.Invoke(func(deps shared.Deps) error {
		var sig = make(chan os.Signal, 1)

		signal.Notify(sig, os.Interrupt)

		go func() {
			deps.Logger.Info("Starting NSQ Consumer")
		}()

		<-sig
		deps.Close()

		return nil
	})

	if err != nil {
		panic(err)
	}
}
