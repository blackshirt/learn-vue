package main

import (
	"context"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"./app/models"
	"./app/types"
)

func main() {
	// setup db handle
	db, err := models.NewDB("./app/types/test.db")
	if err != nil {
		log.Panic(err)
		return
	}
	// the context we want to pass
	// its a good to separate the context key
	ctx := context.WithValue(context.Background(), "db", db)

	// define your root query as entry point
	var schemaConfig = graphql.SchemaConfig{
		Query:    types.RootQuery,
		Mutation: types.MutationsType,
	}

	// define your schema
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Error create schema, error: %v", err)
		return
	}

	// define your graphql handler using graphql-go handler
	gqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	// inject the context to handler
	ctxHandler := customHandler(ctx, gqlHandler)
	// adapts and serve standard http handler
	http.Handle("/gqlhandler", ctxHandler)
	http.ListenAndServe(":8000", nil)
}

// custom middleware to pass context to handler
func customHandler(ctx context.Context, h *handler.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Call context handler to serve HTTP
		h.ContextHandler(ctx, w, r)
	})
}
