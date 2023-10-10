package pubsub

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/BESTSELLER/nightscaler/config"
	"github.com/BESTSELLER/nightscaler/logger"
)

type EventStore interface {
	Listen() error
	Publish(msg string) error
}

func client(projectId string) (*pubsub.Client, error) {
	ctx := context.Background()

	// Create the client.
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		return &pubsub.Client{}, fmt.Errorf("pubsub.NewClient error: %v", err)
	}

	return client, nil
}

func createSubscription(ctx context.Context, client *pubsub.Client) (*pubsub.Subscription, error) {
	// Create subscription name
	subscriptionName := fmt.Sprintf("%s-%s", config.Config.ListenTopic, config.Config.ClusterName)

	var subscription *pubsub.Subscription
	// Check if subscription exists
	subscription = client.Subscription(subscriptionName)
	exists, err := subscription.Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("pubsub.Subscription.Exists error: failed to check if subscription exists %v", err)
	}

	// Create subscription if it doesn't exist
	if !exists {
		subscription, err = client.CreateSubscription(ctx, subscriptionName, pubsub.SubscriptionConfig{
			Topic:       client.Topic(config.Config.ListenTopic),
			Filter:      "attributes.cluster = \"" + config.Config.ClusterName + "\"",
			AckDeadline: 10 * time.Second,
		})
		if err != nil {
			return &pubsub.Subscription{}, fmt.Errorf("pubsub.CreateSubscription error: %v", err)
		}
		logger.Log.Info().Msgf("Subscription %s created to topic %s with attribute filter 'attributes.cluster = %s'", subscriptionName, config.Config.ListenTopic, config.Config.ClusterName)
	} else {
		subConf, err := subscription.Config(ctx)
		if err != nil {
			return nil, fmt.Errorf("pubsub.Subscription.Config error: %v", err)
		}

		logger.Log.Info().Msgf("Subscription %s already exists with attribute filter '%s'", subscriptionName, strings.ReplaceAll(subConf.Filter, "\"", ""))
	}

	return subscription, nil
}
