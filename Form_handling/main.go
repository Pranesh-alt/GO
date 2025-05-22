package main

import (
	"html/template"
	"net/http"
	"strings"
)

func serveForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/form.html"))
	tmpl.Execute(w, nil)
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	name := strings.TrimSpace(r.FormValue("name"))
	email := strings.TrimSpace(r.FormValue("email"))

	// Basic validation
	if name == "" || email == "" {
		http.Error(w, "Name and Email are required fields", http.StatusBadRequest)
		return
	}

	// Respond with success (HTML)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Submission Received</title>
			<style>
				body {
					font-family: 'Segoe UI', sans-serif;
					background-color: #f0f8ff;
					display: flex;
					justify-content: center;
					align-items: center;
					height: 100vh;
				}
				.message {
					text-align: center;
					padding: 40px;
					background: #fff;
					border-radius: 12px;
					box-shadow: 0 6px 20px rgba(0,0,0,0.1);
				}
			</style>
		</head>
		<body>
			<div class="message">
				<h2>Thank you, ` + template.HTMLEscapeString(name) + `!</h2>
				<p>We've received your email: ` + template.HTMLEscapeString(email) + `</p>
			</div>
		</body>
		</html>
	`))
}

func main() {
	// Routes
	http.HandleFunc("/", serveForm)
	http.HandleFunc("/submit", handleSubmit)

	// Static assets if needed later
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	println("Server running at http://localhost:8090")
	http.ListenAndServe(":8090", nil)
}
