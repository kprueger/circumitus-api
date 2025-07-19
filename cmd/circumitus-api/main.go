package main

// import "github.com/joho/godotenv"
// _ = godotenv.Load()

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kprueger/circumitus-api/graph"
	"github.com/kprueger/circumitus-api/graph/generated"
	"github.com/kprueger/circumitus-api/internal/db"
)

func main() {
	fmt.Println("Hello, World!")

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment")
	}

	if os.Getenv("ENV") != "prod" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found - create one by copying .env.example")
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db.Init()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting Waste API server...")
	log.Printf("🚀 Server ready at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
