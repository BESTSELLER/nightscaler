package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/BESTSELLER/nightscaler/config"
	"github.com/BESTSELLER/nightscaler/k8s"
	"github.com/BESTSELLER/nightscaler/logger"
)

type PubSubMsg struct {
	Action    string `json:"action"`
	Namespace string `json:"namespace"`
	Cluster   string `json:"cluster"`
	Duration  int    `json:"duration"`
}

func Listen() error {
	ctx := context.Background()

	// Create the client.
	client, err := client(config.Config.ProjectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient error: %v", err)
	}

	subscription, err := createSubscription(ctx, client)
	if err != nil {
		return err
	}

	// Receive messages
	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		defer msg.Ack()

		logger.Log.Debug().Msgf("Received message: %s", string(msg.Data))

		// Unmarshal message
		var m PubSubMsg
		if err := json.Unmarshal(msg.Data, &m); err != nil {
			logger.Log.Fatal().Err(err).Msg("Failed to unmarshal message")
		}

		if m.Cluster != config.Config.ClusterName {
			return
		}

		if m.Action == "scale_namespace_up" {
			logger.Log.Info().Msgf("Scaling %s namespace %s up", m.Cluster, m.Namespace)
			convertIntToTimeDuration, err := time.ParseDuration(fmt.Sprintf("%dh", m.Duration))
			if err != nil {
				logger.Log.Fatal().Err(err).Msg("Failed to parse duration")
			}

			logger.Log.Debug().Msgf("Duration: %d, Duration in time.Duration: %s", m.Duration, convertIntToTimeDuration)

			k8s.ScaleUp(m.Namespace, convertIntToTimeDuration)
		}

	})
	if err != nil {
		return fmt.Errorf("pubsub.Receive error: %v", err)
	}

	return nil
}
