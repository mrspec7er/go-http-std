package photo

import (
	"github.com/go-chi/chi/v5"
)

func Routes(router chi.Router) {
	controller := &PhotoController{}

	router.Post("/{movieId}", controller.HandlerCreate)
	router.Get("/{movieId}", controller.HandlerGetAll)
	router.Delete("/{movieId}", controller.HandlerDelete)
}
