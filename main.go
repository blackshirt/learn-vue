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

type ContextInjector struct {
	ctx context.Context
	h   http.Handler
}

func (ci *ContextInjector) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ci.h.ServeHTTP(w, r.WithContext(ci.ctx))
}

func main() {
	repo, err := models.NewDB("./app/types/test.db")
	if err != nil {
		log.Panic(err)
	}
	ctx := context.WithValue(context.Background(), "repo", repo)

	// define your root query as entry point
	var schemaConfig = graphql.SchemaConfig{
		Query:    types.RootQuery,
		Mutation: types.MutationsType,
	}
	// define your schema
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Error create schema, error: %v", err)
	}

	// define your graphql handler using graphql-go handler
	gqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// adapts and serve standard http handler
	http.Handle("/gqlhandler", &ContextInjector{ctx, http.HandlerFunc(gqlHandler)})
	http.ListenAndServe(":8000", nil)
}
