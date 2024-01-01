package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mrspec7er/go-http-std/app/repository"
	"golang.org/x/crypto/bcrypt"
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

func (s AuthService) GetUserOauthToken(state string, code string) (*string, error) {
	if state != OauthStateGoogle {
		return nil, errors.New("invalid oauth state")
	}
	token, err := conf.Exchange(context.TODO(), code)
	if err != nil {
		return nil, errors.New("code exchange failed: " + err.Error())
	}

	return &token.AccessToken, nil
}

func (s AuthService) GetUserGoogleInfo(accessToken string) (*UserInfo, error) {
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

func (s AuthService) SaveOauthUser(req *UserInfo) (int, error) {
	status := "INACTIVE"
	if req.VerifiedEmail {
		status = "ACTIVE"
	}

	s.user = repository.User{Name: req.Name, Email: req.Email, Status: status, Role: "USER"}
	err := s.user.Create()
	if err != nil {
		return 500, errors.New("failed reading response body: " + err.Error())
	}

	return 201, nil
}

func (s AuthService) CreateUser(req *repository.User) (int, error) {
	s.user = repository.User{Name: req.Name, Email: req.Email, Password: "UNFILLED", Status: "INACTIVE", Role: "USER"}

	err := s.user.Create()
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s AuthService) GeneratePasswordTokenServices(email string) (*string, error) {
	user, err := s.user.GetByEmail(email)
	if err != nil {
		return  nil, err
	}

	token, err := s.GenerateTokenService(user.Email, 1, "UPDATE_PASSWORD_SECRET")
	if err != nil {
		return  nil, err
	}
	
	return token, nil
}

func (s AuthService) UpdatePasswordService(token string, password string) (int, error) {

	payload, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Failed to parse JWT token!") 
		}

		return []byte("UPDATE_PASSWORD_SECRET"), nil
	})
	if err != nil {
		return 400, err
	}

	claims, ok := payload.Claims.(jwt.MapClaims)

	if ok && payload.Valid {

		encryptedPass, err := bcrypt.GenerateFromPassword([]byte(password), 11)
		if err != nil {
			return  400, err
		}

		email, ok := claims["email"].(string)

		if !ok {
			return  500, errors.New("Failed to convert JWT payload")
		}
		
		s.user.Password = string(encryptedPass)
		s.user.Status = "ACTIVE"

		err = s.user.Update(email)
		if err != nil {
			return  500, err
		}
	}

	return 201, nil
}

func (s AuthService) GenerateTokenService(email string, duration int, secret string) (*string, error)  {

	payload := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp": time.Now().Add(time.Hour * time.Duration(duration)).Unix(),
	})

	token, err := payload.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (s AuthService) LoginService(email string, password string) (*string, *repository.User, error)  {

	user, err := s.user.GetByEmail(email) 
	if err != nil {
		return nil, nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, nil, err
	}

	token, err := s.GenerateTokenService(user.Email, 24, "AUTH_SECRET")
	if err != nil {
		return nil, nil, err
	}

	return token, user, nil
}