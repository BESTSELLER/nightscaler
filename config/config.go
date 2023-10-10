package config

import (
	"github.com/BESTSELLER/nightscaler/instancemetadata"
	"github.com/BESTSELLER/nightscaler/logger"
	"github.com/kelseyhightower/envconfig"
)

const (
	NAMESPACE_ANNOTATION = "downscaler/uptime"
)

type config struct {
	ProjectID    string `envconfig:"PROJECT_ID" required:"true"`
	TimeZone     string `envconfig:"TIMEZONE" default:"UTC"`
	ClusterName  string `envconfig:"CLUSTERNAME" required:"false"`
	ListenTopic  string `envconfig:"LISTEN_TOPIC" required:"true"`
	PublishTopic string `envconfig:"PUBLISH_TOPIC" required:"true"`
	Debug        bool   `envconfig:"DEBUG" default:"false"`
	JsonLogging  bool   `envconfig:"JSON_LOGGING" default:"true"`
}

var Config config

func Init() {
	err := envconfig.Process("NIGHTSCALER", &Config)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to process env variables")
	}

	if Config.ProjectID == "" || Config.ListenTopic == "" || Config.PublishTopic == "" {
		logger.Log.Fatal().Msg("PROJECT_ID, LISTEN_TOPIC and PUBLISH_TOPIC must be set")
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
