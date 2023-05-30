package main

import (
	"github.com/orkarstoft/kscale/pkg/config"
	"github.com/orkarstoft/kscale/pkg/logger"
	"github.com/orkarstoft/kscale/pkg/pubsub"
)

func main() {
	config.Init()
	logger.New(config.Config.Debug, config.Config.JsonLogging)

	logger.Log.Info().Msg("Starting kscale")
	logger.Log.Info().Msgf("Creating subscriber for topic %s in project %s", config.Config.Topic, config.Config.ProjectID)

	err := pubsub.Listen()
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to listen to pubsub")
	}
}
