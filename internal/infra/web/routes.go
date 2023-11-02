package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func(app *Application) routes() http.Handler{
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Heartbeat("/health"))

	mux.Get("/videos", app.ListVideosHandler)
	mux.Get("/videos/{id}", app.GetVideoByIDHandler)
	mux.Get("/video-by-title/{title}", app.GetVideoByTitleHandler)

	return mux
}