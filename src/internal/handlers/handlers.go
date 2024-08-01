package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)
// HomeHandler is the handler for the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("web", "templates", "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}
