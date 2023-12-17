package movie

import "github.com/go-chi/chi/v5"

func Routes(router chi.Router)  {
	handler := &Movie{}
	router.Post("/", handler.Create)
	router.Get("/", handler.GetAll)	
	router.Put("/{id}", handler.Update)
	router.Get("/{id}", handler.GetOne)
	router.Delete("/{id}", handler.Delete)
}