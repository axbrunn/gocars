package web

import (
	"html/template"
	"path/filepath"
)

type TemplateCache map[string]*template.Template

func NewTemplateCashe() (TemplateCache, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseGlob("./templates/layout/base.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./templates/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
