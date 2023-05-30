package pubsublite

import (
	"context"

	"cloud.google.com/go/pubsublite/pscompat"
	"github.com/orkarstoft/kscale/pkg/logger"
)

func client(subscriptionPath string) *pscompat.SubscriberClient {
	ctx := context.Background()

	// Create the subscriber client.
	subscriber, err := pscompat.NewSubscriberClient(ctx, subscriptionPath)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("pscompat.NewSubscriberClientWithSettings error")
	}

	return subscriber
}
