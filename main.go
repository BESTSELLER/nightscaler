package main

import (
	"time"

	"github.com/BESTSELLER/nightscaler/config"
	"github.com/BESTSELLER/nightscaler/k8s"
	"github.com/BESTSELLER/nightscaler/logger"
	"github.com/BESTSELLER/nightscaler/pubsub"
)

func main() {
	config.Init()
	logger.New(config.Config.Debug, config.Config.JsonLogging)

	logger.Log.Info().Msg("Starting nightscaler")
	logger.Log.Info().Msgf("Creating subscriber for topic %s in project %s", config.Config.ListenTopic, config.Config.ProjectID)

	go sendNamespaces()

	err := pubsub.Listen()
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("failed to listen for Pub/Sub messages")
	}

}

func sendNamespaces() {
	for {
		namespaces, err := k8s.List()
		if err != nil {
			logger.Log.Fatal().Err(err).Msg("failed to list namespaces")
		}

		err = pubsub.Publish(namespaces, pubsub.Attributes{
			Entity: "namespace",
			Action: "create",
		})
		if err != nil {
			logger.Log.Fatal().Err(err).Msg("failed to send Pub/Sub message")
		}

		time.Sleep(10 * time.Minute)
	}
}
