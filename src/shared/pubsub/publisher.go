package pubsub

import (
	"context"
	"os"

	pubsub "cloud.google.com/go/pubsub/v2"
)

func NewPublisher(ctx context.Context) (*pubsub.Publisher, error) {
	projectID := os.Getenv("PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return client.Publisher(topicName), nil
}
