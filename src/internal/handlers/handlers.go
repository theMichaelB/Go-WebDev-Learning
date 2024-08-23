package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// HomeHandler is the handler for the home page.
func HomeHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	tmplPath := filepath.Join("web", "templates", "index.html")

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(responseWriter, "Unable to load template", http.StatusInternalServerError)

		return
	}

	err = tmpl.Execute(responseWriter, nil)
	if err != nil {
		http.Error(responseWriter, "Unable to render template", http.StatusInternalServerError)
	}
}
