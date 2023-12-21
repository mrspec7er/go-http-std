package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrspec7er/go-http-std/app/genre"
	"github.com/mrspec7er/go-http-std/app/movie"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("Hello There!"))
	})

	router.Route("/movies", movie.Routes)
	router.Route("/genres", genre.Routes)

	return router
}
