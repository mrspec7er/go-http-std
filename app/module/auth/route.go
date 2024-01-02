package auth

import (
	"github.com/go-chi/chi/v5"
)

func Routes(router chi.Router)  {
	c := &AuthController{}

	router.Get("/login", c.HandleLoginTemplate)

	router.Get("/login/google", c.HandleGoogleLogin)
	router.Get("/callback", c.HandleGoogleAuthCallback)

	router.Post("/send-update-password", c.HandleSendUpdatePassword)
	router.Post("/update-password", c.HandleUpdatePassword)

	router.Post("/login/email", c.HandleEmailLogin)
}