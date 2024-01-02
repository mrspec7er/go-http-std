package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/go-http-std/app/middleware"
)

func Routes(router chi.Router)  {
	c := &UserController{}
	m := &middleware.AuthMiddleware{}

	router.Post("/", c.HandlerCreate)
	router.Get("/", c.HandlerGetAll)	
	router.Get("/{id}", c.HandlerGetOne)
	router.Put("/{id}", HandlerUpdate)
	router.Delete("/{id}", HandlerDelete)
	router.With(m.AuthorizeUser("USER", "STAFF", "ADMIN")).Get("/whoami", c.HandleGetUserInfo)
}