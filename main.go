package main

import (
	"net/http"

	"github.com/eif-courses/fruits/internal/handlers/hello"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", HelloResponse)
	http.ListenAndServe(":3000", r)
}
