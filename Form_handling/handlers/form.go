package handlers

import (
	"Form_handling/config"
	"Form_handling/models"
	"html/template"
	"log"
	"net/http"
	"sync"
)

var csrfTokens = struct {
	sync.RWMutex
	m map[string]bool
}{m: make(map[string]bool)}

func generateCSRFToken() string {
	token := RandString(32)
	csrfTokens.Lock()
	csrfTokens.m[token] = true
	csrfTokens.Unlock()
	return token
}

func RandString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[i%len(letters)]
	}
	return string(b)
}

func FormPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/form.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"CSRFToken": generateCSRFToken(),
	}
	tmpl.Execute(w, data)
}

func SubmitForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Form parse error", http.StatusBadRequest)
		return
	}

	csrfToken := r.FormValue("csrf_token")
	csrfTokens.RLock()
	_, valid := csrfTokens.m[csrfToken]
	csrfTokens.RUnlock()
	if !valid {
		http.Error(w, "Invalid CSRF token", http.StatusForbidden)
		return
	}
	csrfTokens.Lock()
	delete(csrfTokens.m, csrfToken)
	csrfTokens.Unlock()

	contact := models.Contact{
		Name:  r.FormValue("name"),
		Email: r.FormValue("email"),
	}

	errors := contact.Validate()
	if len(errors) > 0 {
		tmpl, _ := template.ParseFiles("templates/form.html")
		data := map[string]interface{}{
			"CSRFToken": generateCSRFToken(),
			"Errors":    errors,
			"Form":      contact,
		}
		tmpl.Execute(w, data)
		return
	}

	result := config.DB.Create(&contact)
	if result.Error != nil {
		log.Println("DB error:", result.Error)
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/success.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, contact)
}
