package pubsublite

import (
	"context"
	"log"

	"cloud.google.com/go/pubsublite/pscompat"
)

func client(subscriptionPath string) *pscompat.SubscriberClient {
	ctx := context.Background()

	// Create the subscriber client.
	subscriber, err := pscompat.NewSubscriberClient(ctx, subscriptionPath)
	if err != nil {
		log.Fatalf("pscompat.NewSubscriberClientWithSettings error: %v", err)
	}

	return subscriber
}
