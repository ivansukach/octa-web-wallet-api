package main

import (
	"github.com/ivansukach/octa-web-wallet-api/repositories/validators"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ivansukach/octa-web-wallet-api/graph"
	"github.com/ivansukach/octa-web-wallet-api/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db, err := sqlx.Connect("postgres",
		"user=su password=su "+
			"host=localhost dbname=mintscan")
	if err != nil {
		log.Fatal(err)
	}
	validatorRps := validators.New(db)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(validatorRps)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
