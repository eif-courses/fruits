package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/eif-courses/fruits/internal/config"
	restapi "github.com/eif-courses/fruits/internal/handlers"
	"github.com/eif-courses/fruits/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
)

func run(connURL string) error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, connURL)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := repository.New(conn)

	// list all authors
	fruits, err := queries.ListFruits(ctx)
	if err != nil {
		return err
	}

	// create an author
	insertedFruit, err := queries.InsertFruit(ctx, repository.InsertFruitParams{
		Name: "Obuolys",
		Colour: "awewea"
	})
	if err != nil {
		return err
	}
	log.Println(insertedFruit)

	log.Println(fruits)

	return nil
}

func main() {

	config.Load()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Now you can use os.Getenv as usual
	if err := run(os.Getenv("DATABASE_URL")); err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", restapi.HelloResponse)

	http.ListenAndServe(":3000", r)
}
