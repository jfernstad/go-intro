package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pressly/chi"
)

// HelloWorld - All exported structs/functions
// should have a comment or the linter will be mad.
type HelloWorld struct {
	Greeting string `json:"verb"` // Object tags
	Receiver string `json:"subject"`
}

// Root [documentation here]
func Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	myObj := HelloWorld{"Hello", "World!"}
	// myObj := HelloWorld{Greeting: "Hello", Receiver: "World!"}

	jsonData, err := json.Marshal(myObj)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // HTTP 500
		fmt.Fprint(w, "{\"error\":\"THAT DID NOT WORK!\"}")
		return
	}

	w.WriteHeader(http.StatusOK) // HTTP 200
	fmt.Fprintf(w, "%s", string(jsonData))
}

func main() {

	// Router implements the `ServeHTTP` method of the `http.Handler` Interface.
	// That is all that is required for an object to be a web service receiver.
	router := chi.NewRouter()

	// Root is of type `http.HandlerFunc`, it will receive all requests on '/'
	router.Get("/", Root)

	// Lambda functions works well in golang
	router.Get("/lambda", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Me be a lambda/closure/anonymous function!")
	})

	// Proudly print the url
	fmt.Println("")
	fmt.Println(",.~<oO( 127.0.0.1:3000 )Oo>~.,")
	fmt.Println("")

	// This opens the socket
	http.ListenAndServe(":3000", router) // Router will receive all requests
}
