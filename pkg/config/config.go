package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/orkarstoft/kscale/pkg/instancemetadata"
	"github.com/orkarstoft/kscale/pkg/logger"
)

type config struct {
	ProjectID   string `envconfig:"PROJECT_ID" required:"true"`
	TimeZone    string `envconfig:"TIMEZONE" default:"UTC"`
	ClusterName string `envconfig:"CLUSTERNAME" required:"false"`
	Topic       string `envconfig:"TOPIC" required:"true"`
	Debug       bool   `envconfig:"DEBUG" default:"false"`
	JsonLogging bool   `envconfig:"JSON_LOGGING" default:"true"`
}

var Config config

func Init() {
	err := envconfig.Process("KSCALE", &Config)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to process env variables")
	}

	if Config.ProjectID == "" || Config.Topic == "" {
		logger.Log.Fatal().Msg("PROJECT_ID and TOPIC must be set")
	}

	if Config.Debug {
		logger.Log.Debug().Msg("Debug logging enabled")
	}

	if Config.JsonLogging && Config.Debug {
		logger.Log.Debug().Msg("JSON logging enabled")
	} else if !Config.JsonLogging && Config.Debug {
		logger.Log.Debug().Msg("JSON logging disabled")
	}

	if Config.ClusterName == "" {
		clustername, err := instancemetadata.GetGCPInstanceMetadata()
		if err != nil {
			logger.Log.Fatal().Err(err).Msg("Failed to get cluster name from instance metadata")
		}
		Config.ClusterName = clustername
	}
}
