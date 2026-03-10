package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"os"
)

func NewPublisher(ctx context.Context) (*pubsub.Topic, error) {
	projectID := os.Getenv("PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return client.Topic(topicName), nil
}
