package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/orkarstoft/kscale/pkg/instancemetadata"
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
		log.Fatal(err.Error())
	}

	if Config.ProjectID == "" || Config.Topic == "" {
		log.Fatal("Missing required env variables")
	}

	if Config.Debug {
		fmt.Printf("[DEBUG]: Config is: %+v\n", Config)
	}

	if Config.ClusterName == "" {
		clustername, err := instancemetadata.GetGCPInstanceMetadata()
		if err != nil {
			log.Fatalf("Failed to get cluster name from GCP Instance Metadata: %v", err)
		}
		Config.ClusterName = clustername
	}
}
