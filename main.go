package main

import (
	"time"

	"github.com/BESTSELLER/nightscaler/config"
	"github.com/BESTSELLER/nightscaler/logger"
	"github.com/BESTSELLER/nightscaler/pubsub"
)

func main() {
	config.Init()
	logger.New(config.Config.Debug, config.Config.JsonLogging)

	logger.Log.Info().Msg("Starting nightscaler")
	logger.Log.Info().Msgf("Creating subscriber for topic %s in project %s", config.Config.ListenTopic, config.Config.ProjectID)

	go func() {
		for {
			pubsub.SendNamespaces()
			time.Sleep(10 * time.Minute)
		}
	}()

	err := pubsub.Listen()
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("failed to listen for Pub/Sub messages")
	}

}
