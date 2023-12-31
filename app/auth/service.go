package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mrspec7er/go-http-std/app/repository"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthService struct {
	user repository.User
}

type UserInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	VerifiedEmail bool `json:"verified_email"`
	Picture string `json:"picture"`
}

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
	token, err := conf.Exchange(context.TODO(), code)
	if err != nil {
		return nil, errors.New("code exchange failed: " + err.Error())
	}

	return &token.AccessToken, nil
}

func (s AuthService) GetUserInfo(accessToken string) (*UserInfo, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, errors.New("failed getting user info: " + err.Error())
	}
	defer response.Body.Close()

	userInfo := &UserInfo{}
	
	if err := json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}

func (s AuthService) SaveUser(req *UserInfo) (int, error) {
	status := "INACTIVE"
	if req.VerifiedEmail {
		status = "ACTIVE"
	}

	user := &repository.User{Name: req.Name, Email: req.Email, Status: status, Role: "USER"}

	err := user.Create()
	if err != nil {
		return 500, errors.New("failed reading response body: " + err.Error())
	}

	return 201, nil
}