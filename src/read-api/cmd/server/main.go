package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	sharedfs "factorio-recipes/shared/firestore"
	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()
	fs, err := sharedfs.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/recipes/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		doc, err := fs.Collection("recipes").Doc(id).Get(r.Context())
		if err != nil {
			http.Error(w, "not found", 404)
			return
		}
		json.NewEncoder(w).Encode(doc.Data())
	})

	log.Println("read-api running on :8080")
	http.ListenAndServe(":8080", r)
}
