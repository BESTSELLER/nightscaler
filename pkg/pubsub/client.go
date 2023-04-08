package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func client(projectId string) (*pubsub.Client, error) {
	ctx := context.Background()

	// Create the client.
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		return &pubsub.Client{}, fmt.Errorf("pubsub.NewClient error: %v", err)
	}

	return client, nil
}
