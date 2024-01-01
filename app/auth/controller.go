package auth

import (
	"fmt"
	"net/http"

	"github.com/mrspec7er/go-http-std/app/utils"
)

type AuthController struct {
	service AuthService
}

const (
	OauthStateGoogle = "google"
)

func (AuthController) HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := conf.AuthCodeURL(OauthStateGoogle)
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
	
	tokenCookie := &http.Cookie{Name: "accessToken", Value: r.FormValue("state") + " " + *token, HttpOnly: false}
	http.SetCookie(w, tokenCookie)

	utils.SuccessMessageResponse(w, "Login Success")
}

func (c *AuthController) HandleGetUserInfo(w http.ResponseWriter, r *http.Request) {
	message := "Authenticated Success"

	user := r.Context().Value("user")

	utils.GetSuccessResponse(w, &message, user, nil)
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