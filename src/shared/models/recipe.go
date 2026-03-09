package models

type Ingredient struct {
	Item   string `json:"item"`
	Amount int    `json:"amount"`
}

type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
	Time        float64      `json:"time"`
	Category    string       `json:"category"`
}
