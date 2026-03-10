package main

import (
	"context"
	"encoding/json"
	"log"

	sharedfs "factorio-recipes/shared/firestore"
	"factorio-recipes/shared/models"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func HandleMessage(ctx context.Context, m PubSubMessage) error {
	fs, err := sharedfs.NewClient(ctx)
	if err != nil {
		return err
	}

	var recipe models.Recipe
	if err := json.Unmarshal(m.Data, &recipe); err != nil {
		return err
	}

	_, err = fs.Collection("recipes").Doc(recipe.Name).Set(ctx, recipe)
	return err
}

func main() {
	log.Println("worker is a Cloud Run Pub/Sub handler, not an HTTP server")
}
