package router

import (
	"net/http"

	"github.com/axbrunn/gocars/internal/config"
	"github.com/axbrunn/gocars/internal/handlers"
	"github.com/julienschmidt/httprouter"
)

func Routes(cfg config.Config) http.Handler {
	router := httprouter.New()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/healthcheck", handlers.HealthcheckHandler(cfg))

	router.HandlerFunc(http.MethodGet, "/", handlers.HomeHandler)

	return router
}
