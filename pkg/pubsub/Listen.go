package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/orkarstoft/kscale/pkg/config"
	"github.com/orkarstoft/kscale/pkg/k8s"
)

func Listen() error {
	ctx := context.Background()

	// Create the client.
	client, err := client(config.Config.ProjectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient error: %v", err)
	}

	// Create subscription
	subscription, err := createSubscription(ctx, client)
	if err != nil {
		return fmt.Errorf("pubsub.CreateSubscription error: %v", err)
	}

	// Receive messages
	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		if config.Config.Debug {
			fmt.Printf("[DEBUG]: Received message: %s", string(msg.Data))
		}

		// Unmarshal message
		var m PubSubMsg
		if err := json.Unmarshal(msg.Data, &m); err != nil {
			panic(err)
		}

		if m.Action == "kscale_scale_namespace_up" {
			fmt.Printf("[INFO]: Scaling %s namespace %s up\n", m.Cluster, m.Namespace)
			convertIntToTimeDuration, err := time.ParseDuration(fmt.Sprintf("%dh", m.Duration))
			if err != nil {
				panic(err)
			}

			if config.Config.Debug {
				fmt.Printf("[DEBUG]: Duration: %d, Duration in time.Duration: %s\n", m.Duration, convertIntToTimeDuration)
			}

			k8s.ScaleNamespaceUp(m.Namespace, convertIntToTimeDuration)
		}

		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("pubsub.Receive error: %v", err)
	}

	return nil
}

type PubSubMsg struct {
	Action    string `json:"action"`
	Namespace string `json:"namespace"`
	Cluster   string `json:"cluster"`
	Duration  int    `json:"duration"`
}
