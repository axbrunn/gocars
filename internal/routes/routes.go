package routes

import (
	"net/http"

	"github.com/axbrunn/gocars/internal/app"
	"github.com/axbrunn/gocars/internal/handlers"
	"github.com/julienschmidt/httprouter"
)

func SetupRoutes(app *app.Application) http.Handler {
	r := httprouter.New()

	fileServer := http.FileServer(http.Dir(app.Config.StaticDir))
	r.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	r.HandlerFunc(http.MethodGet, "/healthcheck", handlers.HealthcheckHandler(*app.Config))

	r.HandlerFunc(http.MethodGet, "/", handlers.HomeHandler)

	return r
}
