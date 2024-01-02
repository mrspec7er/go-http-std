package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrspec7er/go-http-std/app/module/auth"
	"github.com/mrspec7er/go-http-std/app/module/cast"
	"github.com/mrspec7er/go-http-std/app/module/director"
	"github.com/mrspec7er/go-http-std/app/module/genre"
	"github.com/mrspec7er/go-http-std/app/module/movie"
	"github.com/mrspec7er/go-http-std/app/module/user"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("Hello There!"))
	})

	router.Route("/auth", auth.Routes)
	router.Route("/movies", movie.Routes)
	router.Route("/genres", genre.Routes)
	router.Route("/directors", director.Routes)
	router.Route("/casts", cast.Routes)
	router.Route("/users", user.Routes)

	return router
}
