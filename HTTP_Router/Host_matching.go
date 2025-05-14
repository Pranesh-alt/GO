package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func SubdomainHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sub := vars["subdomain"]
	fmt.Fprintf(w, "Subdomain matched: %s\n", sub)
}

func main() {
	r := mux.NewRouter()

	// Host-based route with a dynamic subdomain
	r.Host("{subdomain:[a-z]+}.example.com").
		HandlerFunc(SubdomainHandler)

	fmt.Println("Server started on http//localhost:8082")
	http.ListenAndServe(":8082", r)
}
