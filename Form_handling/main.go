package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("form.html")
	tmpl.Execute(w, nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "ParseForm() error", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	// Print or process form data
	fmt.Fprintf(w, "Received: Name = %s, Email = %s", name, email)
}

func main() {
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/submit", submitHandler)

	fmt.Println("Server started at :8090")
	http.ListenAndServe(":8090", nil)
}
