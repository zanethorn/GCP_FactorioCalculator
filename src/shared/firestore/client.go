package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"os"
)

func NewClient(ctx context.Context) (*firestore.Client, error) {
	projectID := os.Getenv("PROJECT_ID")
	return firestore.NewClient(ctx, projectID)
}
