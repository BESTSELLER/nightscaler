package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	ProjectID   string `envconfig:"PROJECT_ID" required:"true"`
	ClusterName string `envconfig:"CLUSTERNAME" required:"true"`
	Topic       string `envconfig:"TOPIC" required:"true"`
	Debug       bool   `envconfig:"DEBUG" default:"false"`
}

var Config config

func Init() {
	err := envconfig.Process("KSCALE", &Config)
	if err != nil {
		log.Fatal(err.Error())
	}
}
