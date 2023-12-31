package auth

import (
	"context"
	"errors"
	"io"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthService struct {}

var conf *oauth2.Config

func Initialization()  {
	conf = &oauth2.Config{
		ClientID:     "180626421605-3cn0spm34e6851vnp2aintbkibpjg8es.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/auth/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

func (s AuthService) GetUserToken(state string, code string) (*string, error) {
	if state != oauthStateGoogle {
		return nil, errors.New("invalid oauth state")
	}
	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		return nil, errors.New("code exchange failed: " + err.Error())
	}

	return &token.AccessToken, nil
}

func (s AuthService) GetUserInfo(accessToken string) ([]byte, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, errors.New("failed getting user info: " + err.Error())
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("failed reading response body: " + err.Error())
	}
	return contents, nil
}