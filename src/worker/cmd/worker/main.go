package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	sharedfs "factorio-recipes/shared/firestore"
	"factorio-recipes/shared/models"
)

func handleMessage(ctx context.Context, m *pubsub.Message) {
	fs, err := sharedfs.NewClient(ctx)
	if err != nil {
		log.Println("firestore error:", err)
		m.Nack()
		return
	}

	var recipe models.Recipe
	if err := json.Unmarshal(m.Data, &recipe); err != nil {
		log.Println("json error:", err)
		m.Nack()
		return
	}

	_, err = fs.Collection("recipes").Doc(recipe.Name).Set(ctx, recipe)
	if err != nil {
		log.Println("firestore write error:", err)
		m.Nack()
		return
	}

	m.Ack()
}

func main() {
	ctx := context.Background()

	projectID := os.Getenv("PROJECT_ID")
	subName := "recipe-writes-sub"

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}

	sub := client.Subscription(subName)

	log.Println("Worker listening for messages on:", subName)

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		handleMessage(ctx, msg)
	})

	if err != nil {
		log.Fatal(err)
	}
}
