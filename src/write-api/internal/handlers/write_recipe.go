func (h *Handler) WriteRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	json.NewDecoder(r.Body).Decode(&recipe)

	if err := validation.ValidateRecipe(recipe); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	data, _ := json.Marshal(recipe)
	h.Topic.Publish(r.Context(), &pubsub.Message{Data: data})

	w.WriteHeader(http.StatusAccepted)
}

