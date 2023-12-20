package movie

import (
	"github.com/go-chi/chi/v5"
)

func Routes(router chi.Router)  {
	controller := &MovieController{}

	router.Post("/", controller.HandlerCreate)
	router.Get("/", controller.HandlerGetAll)	
	router.Put("/{id}", HandlerUpdate)
	router.Get("/{id}", HandlerGetOne)
	router.Delete("/{id}", HandlerDelete)
}