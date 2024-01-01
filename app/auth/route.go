package auth

import (
	"github.com/go-chi/chi/v5"
)

func Routes(router chi.Router)  {
	controller := &AuthController{}
	middleware := &AuthMiddleware{}

	router.Get("/login", controller.HandleLoginTemplate)

	router.Get("/login/google", controller.HandleGoogleLogin)
	router.Get("/callback", controller.HandleGoogleAuthCallback)
	router.With(middleware.AuthorizeUser("USER")).Get("/whoami", controller.HandleGetUserInfo)

	router.Post("/send-update-password", controller.HandleSendUpdatePassword)
	router.Post("/update-password", controller.HandleUpdatePassword)

	router.Post("/login/email", controller.HandleEmailLogin)
}