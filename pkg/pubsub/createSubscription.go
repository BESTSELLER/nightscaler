package pubsub

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/orkarstoft/kscale/pkg/config"
	"github.com/orkarstoft/kscale/pkg/random"
)

func createSubscription(ctx context.Context, client *pubsub.Client) (pubsub.Subscription, error) {
	// Create subscription
	subName := fmt.Sprintf("kscale-%s-%s", config.Config.ClusterName, random.String(5))
	subscription, err := client.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{
		Topic:       client.Topic(config.Config.Topic),
		Filter:      "attributes.cluster = \"" + config.Config.ClusterName + "\"",
		AckDeadline: 10 * time.Second,
	})
	if err != nil {
		return pubsub.Subscription{}, fmt.Errorf("pubsub.CreateSubscription error: %v", err)
	}

	return *subscription, nil
}
