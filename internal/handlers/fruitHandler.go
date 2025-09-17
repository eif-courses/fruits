package handlers

import (
	"encoding/json"
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
