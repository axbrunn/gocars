package handlers

import (
	"html/template"
	"log/slog"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"internal/templates/layout/base.tmpl",
		"internal/templates/layout/nav.tmpl",
		"internal/templates/public/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		slog.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		slog.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
