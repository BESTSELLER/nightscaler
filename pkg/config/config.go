package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	ProjectID      string `envconfig:"PROJECT_ID"`
	Zone           string `envconfig:"ZONE"`
	ClusterName    string `envconfig:"CLUSTERNAME"`
	SubscriptionID string `envconfig:"SUBSCRIPTION_ID"`
}

var Config config

func Init() {
	err := envconfig.Process("KSCALE", &Config)
	if err != nil {
		log.Fatal(err.Error())
	}

	format := "PROJECT_ID: %s\nZONE: %s\nCLUSTERNAME: %s\nSUBSCRIPTION_ID: %s\n"
	_, err = fmt.Printf(format, Config.ProjectID, Config.Zone, Config.ClusterName, Config.SubscriptionID)
	if err != nil {
		log.Fatal(err.Error())
	}
}
