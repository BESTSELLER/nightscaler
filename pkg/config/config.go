package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	ProjectID   string `envconfig:"PROJECT_ID"`
	ClusterName string `envconfig:"CLUSTERNAME"`
	Topic       string `envconfig:"TOPIC"`
}

var Config config

func Init() {
	err := envconfig.Process("KSCALE", &Config)
	if err != nil {
		log.Fatal(err.Error())
	}

	if Config.ProjectID == "" {
		log.Fatal("PROJECT_ID is not set")
	}

	if Config.ClusterName == "" {
		log.Fatal("CLUSTERNAME is not set")
	}

	if Config.Topic == "" {
		log.Fatal("TOPIC is not set")
	}
}
