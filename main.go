package main

import (
	"fmt"

	"github.com/orkarstoft/kscale/pkg/config"
	"github.com/orkarstoft/kscale/pkg/pubsublite"
)

var ClusterName string

func main() {
	config.Init()
	fmt.Printf("projectID: %s, zone: %s, subscriptionID: %s\n", config.Config.ProjectID, config.Config.Zone, config.Config.SubscriptionID)
	subscriptionPath := fmt.Sprintf("projects/%s/locations/%s/subscriptions/%s", config.Config.ProjectID, config.Config.Zone, config.Config.SubscriptionID)

	// Listen to pub sub lite messages
	pubsublite.Listen(subscriptionPath)
}
