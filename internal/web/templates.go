package web

import (
	"html/template"
	"path/filepath"
)

type TemplateCache map[string]*template.Template

type ThemeCache map[string]TemplateCache

func NewTemplateCache() (ThemeCache, error) {
	cache := ThemeCache{}

	siteCache, err := buildCache("./templates/site")
	if err != nil {
		return nil, err
	}
	cache["site"] = siteCache

	themes, err := filepath.Glob("./templates/tenants/*")
	if err != nil {
		return nil, err
	}

	for _, theme := range themes {
		slug := filepath.Base(theme)
		themeCache, err := buildCache(theme)
		if err != nil {
			return nil, err
		}
		cache[slug] = themeCache
	}

	return cache, nil
}

func buildCache(dir string) (TemplateCache, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "pages", "*.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseGlob(filepath.Join(dir, "base.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "partials", "*.tmpl"))
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
