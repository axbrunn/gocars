package web

import (
	"net/http"
	"time"
)

type TemplateData struct {
	CurrentYear int
	Title       string
}

func NewTemplateData(r *http.Request) TemplateData {
	return TemplateData{
		CurrentYear: time.Now().Year(),
	}
}
