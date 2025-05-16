package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"strings"
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

	// Subroute for product-related routes
	productRouter := r.PathPrefix("/products/").Subrouter()

	productRouter.HandleFunc("/{id}", ProductDetailHandler)

	// Middleware to log the request method and URL
	type MiddlewareFunc func(http.Handler) http.Handler

	// Define the directory path for static files
	dir := "./static"
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir(dir))),
	)

	// Define a route with a named variable
	r.HandleFunc("/articles/{category}", ArticlesCategoryHandler).
		Schemes("https").
		Headers("Content-Type", "application/json", "Content-Type", "application/text").
		Methods("GET", "POST").
		Queries("type", "premium", "sort", "asc"). //filter?type=premium&sort=asc
		Name("article")

	// Build a URL Dynamically
	url, err := r.Get("article").URL("category", "technology", "id", "42")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(url.String())

	// Query Parameters

	q := url.Query()
	q.Add("sort", "desc")
	url.RawQuery = q.Encode()

	fmt.Println(url.String()) // /articles/technology/42?sort=desc

	// custom Usage Idea: List Routes on /debug/routes
	r.HandleFunc("/debug/routes", func(w http.ResponseWriter, req *http.Request) {
		r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			pathTemplate, _ := route.GetPathTemplate()
			methods, _ := route.GetMethods()
			fmt.Fprintf(w, "%s %v\n", strings.Join(methods, ","), pathTemplate) //output : GET,POST /articles/{category}
			return nil
		})
	})

	// CORS Middleware
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(r)

	fmt.Println("Server started on http//localhost:8081")

	http.ListenAndServe(":8081", handler)

}
