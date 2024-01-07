package cast

import (
	"github.com/go-chi/chi/v5"
)

func Routes(router chi.Router) {
	controller := &CastController{}

	router.Post("/", controller.HandlerCreate)
	router.Get("/", controller.HandlerGetAll)
	router.Get("/{id}", controller.HandlerGetOne)
	router.Put("/{id}", HandlerUpdate)
	router.Delete("/{id}", HandlerDelete)
}
