package movie

import (
	"github.com/go-chi/chi/v5"
)

func Routes(router chi.Router) {
	controller := &MovieController{}

	router.Post("/", controller.HandlerCreate)
	router.Get("/", controller.HandlerGetAll)
	router.Get("/{id}", controller.HandlerGetOne)
	router.Put("/", controller.HandlerUpdate)
	router.Put("/thumbnail", controller.HandlerUpdateThumbnail)
	router.Delete("/{id}", HandlerDelete)
}
