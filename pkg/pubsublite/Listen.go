package pubsublite

import (
	"context"
	"sync/atomic"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/pubsublite/pscompat"
	"github.com/orkarstoft/kscale/pkg/config"
	"github.com/orkarstoft/kscale/pkg/k8s"
	"github.com/orkarstoft/kscale/pkg/logger"
)

func Listen(subscriptionPath string) {
	ctx := context.Background()
	subscriber := client(subscriptionPath)

	// Listen for messages
	logger.Log.Info().Msgf("Listening to messages on %s", subscriptionPath)
	var receiveCount int32

	cctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Receive blocks until the context is cancelled or an error occurs.
	if err := subscriber.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		logger.Log.Debug().Msgf("Received message: %s", string(msg.Data))
		atomic.AddInt32(&receiveCount, 1)

		_, err := pscompat.ParseMessageMetadata(msg.ID)
		if err != nil {
			logger.Log.Fatal().Err(err).Msgf("Failed to parse %q", msg.ID)
		}

		if string(msg.Data) == "kscale-scale-namespace-up" {
			logger.Log.Info().Msgf("Scaling %s namespace %s up", msg.Attributes["clustername"], msg.Attributes["namespace"])
			if msg.Attributes["clustername"] == config.Config.ClusterName {
				k8s.ScaleNamespaceUp(msg.Attributes["namespace"], 1)
			} else {
				logger.Log.Warn().Msgf("Received %s - not my cluster", string(msg.Data))
			}
			msg.Ack()
		} else {
			logger.Log.Warn().Msgf("Received %s - unknown message", string(msg.Data))
			msg.Ack()
		}
	}); err != nil {
		logger.Log.Fatal().Err(err).Msg("SubscriberClient.Receive error")
	}

	logger.Log.Info().Msgf("Received %d messages", receiveCount)
}
