package auth

import (
	"fmt"
	"net/http"

	"github.com/mrspec7er/go-http-std/app/utils"
)

type AuthMiddleware struct {
	service AuthService
}

func (m AuthMiddleware) AuthenticatedUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("accessToken")
		if err != nil {
			utils.InternalServerErrorHandler(w, 500, err)
			return
		}
	
		info, err := m.service.GetUserInfo(cookie.Value)
		if err != nil {
			utils.InternalServerErrorHandler(w, 500, err)
			return
		}

		// TODO: Save user info to context
		fmt.Println("USER_INFO", info)
	
		next.ServeHTTP(w, r)
	})
}