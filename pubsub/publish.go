package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/BESTSELLER/nightscaler/config"
	"github.com/BESTSELLER/nightscaler/logger"
	v1 "k8s.io/api/core/v1"
)

type Attributes struct {
	Entity string `json:"entity"`
	Action string `json:"action"`
}

func Publish(msg []v1.Namespace, att Attributes) error {
	ctx := context.Background()

	// Create the client.
	client, err := client(config.Config.ProjectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient error: %v", err)
	}

	pubsubPublisher := client.Topic(config.Config.PublishTopic)
	topicExists, err := pubsubPublisher.Exists(ctx)
	if err != nil {
		return err
	}
	if !topicExists {
		return fmt.Errorf("topic '%v' was not found", config.Config.PublishTopic)
	}
	defer pubsubPublisher.Stop()

	msgJSON, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("error marshaling entity: %v", err)
	}

	// Send a message containing the message and attributes
	msgId, err := pubsubPublisher.Publish(ctx, &pubsub.Message{
		Data: []byte(msgJSON),
		Attributes: map[string]string{
			"entity":  att.Entity,
			"action":  att.Action,
			"cluster": config.Config.ClusterName,
		},
	}).Get(ctx)
	if err != nil {
		return err
	}

	logger.Log.Debug().Msgf("Published a message; msg ID: %v", msgId)

	return nil
}
