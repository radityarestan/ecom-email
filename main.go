package main

import (
	"github.com/radityarestan/ecom-email/internal/di"
	"log"
	"os"
	"os/signal"
)

func main() {
	deps, err := di.InitDeps()
	if err != nil {
		log.Fatalf("[FATAL] Failed to inject dependency: %v", err)
		return
	}

	go func() {
		log.Println("[INFO] Starting Ecommerce Service Consumer")
		if err := deps.NSQConsumer.Start(deps.Config.Sender.Email, deps.Config.Sender.Password); err != nil {
			log.Fatalf("[FATAL] Failed to start NSQ Consumer: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	deps.NSQConsumer.Stop()
	log.Println("[INFO] Ecommerce Email Service Consumer Stopped")
}
