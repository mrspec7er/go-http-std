package auth

import (
	"context"
	"net/http"
	"strings"

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

		token := strings.Split(cookie.Value, " ")
	
		info, err := m.GetUserInfo(token[0], token[1])
		if err != nil {
			utils.InternalServerErrorHandler(w, 500, err)
			return
		}

		ctx := context.WithValue(r.Context(), "user", &info)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m AuthMiddleware) GetUserInfo(bearer string, accessToken string) (*UserInfo, error) {
	var info *UserInfo
	var err error

	if bearer == OauthStateGoogle {
		info, err = m.service.GetUserGoogleInfo(accessToken)
		if err != nil {
			return info, err
		}
	}

	return info, nil
}