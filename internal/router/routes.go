package router

import (
	"net/http"

	"github.com/axbrunn/gocars/internal/application"
	"github.com/axbrunn/gocars/internal/handlers"
	"github.com/julienschmidt/httprouter"
)

func Routes(app application.Application) http.Handler {
	router := httprouter.New()

	fileServer := http.FileServer(http.Dir(app.Config.StaticDir))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/healthcheck", handlers.HealthcheckHandler(*app.Config))

	router.HandlerFunc(http.MethodGet, "/", handlers.HomeHandler)

	return router
}
