package main

import (
	"fmt"

	"github.com/orkarstoft/kscale/pkg/config"
	"github.com/orkarstoft/kscale/pkg/pubsub"
)

func main() {
	config.Init()

	fmt.Println("Starting KScale")
	fmt.Printf("Creating subscriber for topic %s in project %s\n", config.Config.Topic, config.Config.ProjectID)
	err := pubsub.Listen()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}
