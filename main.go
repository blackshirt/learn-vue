package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"./app/types"
)

func main() {
	// define your root query as entry point
	var schemaConfig = graphql.SchemaConfig{
		Query:    types.RootQuery,
		Mutation: types.MutationsType,
	}
	// define your schema
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Erro create schema, error: %v", err)
	}

	// define your graphql handler using graphql-go handler
	gqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// adapts and serve standard http handler
	http.Handle("/gqlhandler", gqlHandler)
	http.ListenAndServe(":8000", nil)
}
