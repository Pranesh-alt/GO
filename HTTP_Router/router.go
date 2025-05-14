package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Handler to extract "category" from the route
func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)          // Extracts path variables into a map
	category := vars["category"] // Gets the value of the "category" variable
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", category)
}

func main() {
	r := mux.NewRouter()

	// Define a route with a named variable
	r.HandleFunc("/articles/{category}", ArticlesCategoryHandler)
	fmt.Println("Server started on http//localhost:8080")

	http.ListenAndServe(":8080", r)

}
