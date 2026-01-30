package web

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/axbrunn/gocars/internal/http/respond"
)

type Renderer struct {
	TemplateCache TemplateCache
}

func NewRenderer(tc TemplateCache) *Renderer {
	return &Renderer{
		TemplateCache: tc,
	}
}

func (rendr *Renderer) Render(w http.ResponseWriter, r *http.Request, status int, page string, td TemplateData) {
	ts, ok := rendr.TemplateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		respond.ServerError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)

	// Write the template to the buffer, instead of straight to the http.ResponseWriter.
	err := ts.ExecuteTemplate(buf, "base", td)
	if err != nil {
		respond.ServerError(w, r, err)
		return
	}

	// If the template is written to the buffer without any errors, it's safe
	// to go ahead and write the HTTP status code to http.ResponseWriter.
	w.WriteHeader(status)

	// Write the contents of the buffer to the http.ResponseWriter.
	buf.WriteTo(w)
}
