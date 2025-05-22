package main

import (
	"Form_handling/config"
	"Form_handling/handlers"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.FormPage)
	http.HandleFunc("/submit", handlers.SubmitForm)

	log.Println("Server running at http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
