package web

import (
	"html/template"
	"path/filepath"
)

type TemplateCache map[string]*template.Template

func NewTemplateCashe() (TemplateCache, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/public/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		files := []string{
			"./templates/layout/base.tmpl",
			"./templates/layout/nav.tmpl",
			page,
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
