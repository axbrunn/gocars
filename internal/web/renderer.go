package web

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/axbrunn/gocars/internal/http/middleware"
	"github.com/axbrunn/gocars/internal/http/respond"
)

type Renderer struct {
	ThemeCache ThemeCache
}

func NewRenderer(tc ThemeCache) *Renderer {
	return &Renderer{
		ThemeCache: tc,
	}
}

func (rendr *Renderer) Render(w http.ResponseWriter, r *http.Request, status int, page string, td TemplateData) {
	var theme string

	if tenant, ok := middleware.TenantFromContext(r.Context()); ok {
		td.Title = tenant.Name
		theme = tenant.TemplateSlug
	} else {
		td.Title = "GoCars"
		theme = "site"
	}

	tc, ok := rendr.ThemeCache[theme]
	if !ok {
		err := fmt.Errorf("the theme %s does not exist", theme)
		respond.ServerError(w, r, err)
		return
	}

	ts, ok := tc[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist in theme %s", page, theme)
		respond.ServerError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", td)
	if err != nil {
		respond.ServerError(w, r, err)
		return
	}

	w.WriteHeader(status)
	buf.WriteTo(w)
}
