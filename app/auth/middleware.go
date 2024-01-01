package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mrspec7er/go-http-std/app/repository"
	"github.com/mrspec7er/go-http-std/app/utils"
)

type AuthMiddleware struct {
	service AuthService
}

func (m AuthMiddleware) AuthenticatedUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("accessToken")
		if err != nil {
			utils.InternalServerErrorHandler(w, 400, err)
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

func (m AuthMiddleware) GetUserInfo(bearer string, accessToken string) (*repository.User, error) {

	if bearer == DefaultAuth {
		payload, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("Failed to parse JWT token!")
			}
			return []byte("AUTH_SECRET"), nil
		})
	
		if err != nil {
			return nil, err
		}
	
		claims, ok := payload.Claims.(jwt.MapClaims)
		if !ok || !payload.Valid {
			return nil, errors.New("Failed to encoded token payload")
		}

		email, ok := claims["email"].(string)

		if !ok {
			fmt.Println("TYPE: ", claims)
			return nil, errors.New("Failed to parse token payload")
		}

		user, err := m.service.user.GetByEmail(email)
		if err != nil {
			return nil, err
		} 

		return user, nil
	}

	if bearer == OauthStateGoogle {
		result, err := m.service.GetUserGoogleInfo(accessToken)
		if err != nil {
			return nil, err
		}

		user, err := m.service.user.GetByEmail(result.Email)
		if err != nil {
			return nil, err
		} 

		return user, nil
	}

	return nil, errors.New("Failed Token Bearer")

}