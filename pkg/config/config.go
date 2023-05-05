package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/orkarstoft/kscale/pkg/instancemetadata"
)

type config struct {
	ProjectID   string `envconfig:"PROJECT_ID" required:"true"`
	ClusterName string `envconfig:"CLUSTERNAME" required:"false"`
	Topic       string `envconfig:"TOPIC" required:"true"`
	Debug       bool   `envconfig:"DEBUG" default:"false"`
}

var Config config

func Init() {
	err := envconfig.Process("KSCALE", &Config)
	if err != nil {
		log.Fatal(err.Error())
	}

	if Config.ProjectID == "" || Config.Topic == "" {
		log.Fatal("Missing required env variables")
	}

	if Config.ClusterName == "" {
		clustername, err := instancemetadata.GetGCPInstanceMetadata()
		if err != nil {
			log.Fatalf("Failed to get cluster name from GCP Instance Metadata: %v", err)
		}
		Config.ClusterName = clustername
	}
}
