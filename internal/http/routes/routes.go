package routes

import (
	"net/http"

	"github.com/axbrunn/gocars/internal/app"
	"github.com/axbrunn/gocars/internal/http/handlers"
	"github.com/julienschmidt/httprouter"
)

func SetupRoutes(app *app.Application) http.Handler {
	r := httprouter.New()

	fileServer := http.FileServer(http.Dir(app.Config.StaticDir))
	r.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	// handlers
	handleHealth := handlers.NewHealthcheckHandler(app.Logger, app.Config)
	handleHome := handlers.NewHomeHandler(app.Logger, app.Renderer)

	// end points
	r.HandlerFunc(http.MethodGet, "/healthcheck", handleHealth.Check)
	r.HandlerFunc(http.MethodGet, "/", handleHome.Index)

	return r
}
