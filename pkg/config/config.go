package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/orkarstoft/kscale/pkg/instancemetadata"
	"github.com/rs/zerolog/log"
)

type config struct {
	ProjectID   string `envconfig:"PROJECT_ID" required:"true"`
	TimeZone    string `envconfig:"TIMEZONE" default:"Europe/Copenhagen"`
	ClusterName string `envconfig:"CLUSTERNAME" required:"false"`
	Topic       string `envconfig:"TOPIC" required:"true"`
	Debug       bool   `envconfig:"DEBUG" default:"false"`
}

var Config config

func Init() {
	err := envconfig.Process("KSCALE", &Config)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to process env variables")
	}

	if Config.ProjectID == "" || Config.Topic == "" {
		log.Fatal().Msg("PROJECT_ID and TOPIC must be set")
	}

	if Config.Debug {
		log.Debug().Msg("Debug logging enabled")
	}

	if Config.ClusterName == "" {
		clustername, err := instancemetadata.GetGCPInstanceMetadata()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get cluster name from instance metadata")
		}
		Config.ClusterName = clustername
	}
}
