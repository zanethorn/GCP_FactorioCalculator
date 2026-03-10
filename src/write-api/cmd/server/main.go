package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	pubsub "cloud.google.com/go/pubsub/v2"
	"factorio-recipes/shared/models"
	sharedpub "factorio-recipes/shared/pubsub"
	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()
	publisher, err := sharedpub.NewPublisher(ctx)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/recipes/{id}", func(w http.ResponseWriter, r *http.Request) {
		var recipe models.Recipe
		json.NewDecoder(r.Body).Decode(&recipe)

		data, _ := json.Marshal(recipe)
		publisher.Publish(r.Context(), &pubsub.Message{Data: data})

		w.WriteHeader(http.StatusAccepted)
	}).Methods("PUT")

	log.Println("write-api running on :8080")
	http.ListenAndServe(":8080", r)
}
