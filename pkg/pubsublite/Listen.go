package pubsublite

import (
	"context"
	"fmt"
	"log"
	"sync/atomic"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/pubsublite/pscompat"
	"github.com/orkarstoft/kscale/pkg/config"
	"github.com/orkarstoft/kscale/pkg/k8s"
)

func Listen(subscriptionPath string) {
	ctx := context.Background()
	subscriber := client(subscriptionPath)

	// Listen for messages
	log.Printf("Listening to messages on %s", subscriptionPath)
	var receiveCount int32

	cctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Receive blocks until the context is cancelled or an error occurs.
	if err := subscriber.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Println("message received")
		atomic.AddInt32(&receiveCount, 1)

		_, err := pscompat.ParseMessageMetadata(msg.ID)
		if err != nil {
			log.Fatalf("Failed to parse %q: %v", msg.ID, err)
		}

		if string(msg.Data) == "kscale-scale-namespace-up" {
			fmt.Printf("Scaling %s namespace %s up\n", msg.Attributes["clustername"], msg.Attributes["namespace"])
			if msg.Attributes["clustername"] == config.Config.ClusterName {
				k8s.ScaleNamespaceUp(msg.Attributes["namespace"], 1)
			} else {
				fmt.Printf("Received %s - not my cluster\n", string(msg.Data))
			}
			msg.Ack()
		} else {
			fmt.Printf("Received %s - unknown message\n", string(msg.Data))
			msg.Ack()
		}
	}); err != nil {
		log.Fatalf("SubscriberClient.Receive error: %v", err)
	}

	fmt.Printf("Received %d messages\n", receiveCount)
}
