package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// define schema, with our rootQuery and rootMutation
// TODO root query and mutation are pending
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: packMutation,
})

// HttpGet is the handler function to attend all http get requests
// It gets the graphql parameter and delegates to business logic
// function.
func httpGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()["query"][0]
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), 400)
		log.Println(err.Error())
	}
}

// httpPost is the handler function to attend all http post requests
// It gets the http post parameters and delegates to graphql schema.
func httpPost(w http.ResponseWriter, r *http.Request) {
	// parse http.Request into handler.RequestOptions
	opts := handler.NewRequestOptions(r)

	// inject context objects http.ResponseWrite and *http.Request into rootValue
	// there is an alternative example of using `net/context` to store context
	// instead of using rootValue
	rootValue := map[string]interface{}{
		"response": w,
		"request":  r,
	}

	// execute graphql query
	// here, we passed in Query, Variables and OperationName extracted from http.Request
	params := graphql.Params{
		Schema:         schema,
		RequestString:  opts.Query,
		VariableValues: opts.Variables,
		OperationName:  opts.OperationName,
		RootObject:     rootValue,
	}

	result := graphql.Do(params)

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}

	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), 400)
		log.Println(err.Error())
	}
}
