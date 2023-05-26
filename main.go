package main

import (
	"github.com/orkarstoft/kscale/pkg/config"
	"github.com/orkarstoft/kscale/pkg/logging"
	"github.com/orkarstoft/kscale/pkg/pubsub"
	"github.com/rs/zerolog/log"
)

func main() {
	config.Init()
	logging.Init()

	log.Info().Msg("Starting kscale")
	log.Info().Msgf("Creating subscriber for topic %s in project %s", config.Config.Topic, config.Config.ProjectID)

	err := pubsub.Listen()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen to pubsub")
	}
}
