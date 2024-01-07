package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mrspec7er/go-http-std/app/utils"
)

type AuthController struct {
	service AuthService
}

type AuthPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Token           string `json:"token"`
}

func (AuthController) HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := conf.AuthCodeURL(utils.OauthStateGoogle)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (c *AuthController) HandleGoogleAuthCallback(w http.ResponseWriter, r *http.Request) {
	token, err := c.service.GetUserOauthToken(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		utils.InternalServerErrorHandler(w, 500, err)
		return
	}

	info, err := c.service.GetUserGoogleInfo(*token)
	if err != nil {
		utils.InternalServerErrorHandler(w, 500, err)
		return
	}

	status, err := c.service.SaveOauthUser(info)
	if err != nil {
		utils.InternalServerErrorHandler(w, status, err)
		return
	}

	tokenCookie := &http.Cookie{Name: "accessToken", Value: r.FormValue("state") + " " + *token, HttpOnly: false, Path: "/"}
	http.SetCookie(w, tokenCookie)

	utils.SuccessMessageResponse(w, "Login Success")
}

func (c *AuthController) HandleSendUpdatePassword(w http.ResponseWriter, r *http.Request) {
	var payload AuthPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.BadRequestHandler(w)
		return
	}

	token, err := c.service.GeneratePasswordTokenServices(payload.Email)
	if err != nil {
		utils.BadRequestHandler(w)
		return
	}

	fmt.Println("USER_TOKEN", *token)

	// utils.SendUpdatePassword(token)

	utils.SuccessMessageResponse(w, "Update Password Url sended to: "+payload.Email)
}

func (c *AuthController) HandleUpdatePassword(w http.ResponseWriter, r *http.Request) {
	var payload AuthPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.BadRequestHandler(w)
		return
	}

	if payload.Password != payload.ConfirmPassword {
		utils.BadRequestHandler(w)
		return
	}

	status, err := c.service.UpdatePasswordService(payload.Token, payload.Password)
	if err != nil {
		utils.InternalServerErrorHandler(w, status, err)
		return
	}

	utils.SuccessMessageResponse(w, "Password Updated")
}

func (c *AuthController) HandleLoginTemplate(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `
		<html>
			<body>
				<a href="/auth/login/google">Google Log In</a>
			</body>
		</html>`
	fmt.Fprintf(w, htmlIndex)
}

func (c *AuthController) HandleEmailLogin(w http.ResponseWriter, r *http.Request) {
	var payload AuthPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.BadRequestHandler(w)
		return
	}

	token, user, err := c.service.LoginService(payload.Email, payload.Password)
	if err != nil {
		utils.BadRequestHandler(w)
		return
	}

	tokenCookie := &http.Cookie{Name: "accessToken", Value: utils.DefaultAuth + " " + *token, HttpOnly: false, Path: "/"}
	http.SetCookie(w, tokenCookie)

	utils.GetSuccessResponse(w, nil, user, nil)
}
