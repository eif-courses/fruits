package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eif-courses/fruits/internal/services"
)

type FruitHandler struct {
	fruitService *services.FruitService
}

func NewFruitHandler(f *services.FruitService) *FruitHandler {
	return &FruitHandler{
		fruitService: f,
	}
}

func (f *FruitHandler) GetFruits(w http.ResponseWriter, r *http.Request) {

	fruits, err := f.fruitService.GetFruits(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(fruits)

}

func (f *FruitHandler) InsertFruit(w http.ResponseWriter, r *http.Request) {

	type InsertFruitRequest struct {
		Name   string `json:"name"`
		Colour string `json:"colour"`
	}
	var req InsertFruitRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}
	if req.Colour == "" {
		http.Error(w, "colour is required", http.StatusBadRequest)
		return
	}

	fruit, err := f.fruitService.InsertFruit(r.Context(), req.Name, req.Colour)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create fruit: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the created fruit
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fruit)

}
