package main

import (
	"net/http"

	"github.com/eif-courses/fruits/internal/config"
	restapi "github.com/eif-courses/fruits/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	config.Load()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", restapi.HelloResponse)

	http.ListenAndServe(":3000", r)
}
