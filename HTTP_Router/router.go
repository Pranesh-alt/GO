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

func ProductDetailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Extracts path variables into a map
	id := vars["id"]    // Gets the value of the "id" variable
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product ID: %v\n", id)
}

func main() {
	r := mux.NewRouter()

	// Subrouter for product-related routes
	productRouter := r.PathPrefix("/products/").Subrouter()

	productRouter.HandleFunc("/{id}", ProductDetailHandler)

	// Define a route with a named variable
	r.HandleFunc("/articles/{category}", ArticlesCategoryHandler).
		Schemes("https").
		Headers("X-Requested-With", "XMLHttpRequest").
		Methods("GET", "POST").
		Queries("type", "premium", "sort", "asc") //filter?type=premium&sort=asc

	fmt.Println("Server started on http//localhost:8081")

	http.ListenAndServe(":8081", r)

}
