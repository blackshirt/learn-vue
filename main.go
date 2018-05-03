package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	// type ObjectConfig struct {
	// Name        string      `json:"name"`
	// Interfaces  interface{} `json:"interfaces"`
	// Fields      interface{} `json:"fields"`
	// IsTypeOf    IsTypeOfFn  `json:"isTypeOf"`
	// Description string      `json:"description"`
	// }
	var rootObjectConfig = graphql.ObjectConfig{
		Name
	}
	var rootQueryType = graphql.NewObject(rootObjectConfig)
	// type SchemaConfig struct {
	// Query        *Object
	// Mutation     *Object
	// Subscription *Object
	// Types        []Type
	// Directives   []*Directive
	//}
	var schemaConfig = graphql.SchemaConfig{
		Query: rootQueryType,
	}
	// define your schema
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Erro create schema")
	}

	// define your graphql handler using graphql-go handler
	gqlHandler := handler.New(&handler.Config{
		Schema:   schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// adapts and serve standard http handler
	http.Handle("/gqlhandler", gqlHandler)
	http.ListenAndServe(":5000", nil)
}
