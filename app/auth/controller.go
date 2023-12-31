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
	oauthStateGoogle = "google"
)

func (AuthController) HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := conf.AuthCodeURL(oauthStateGoogle)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (c AuthController) HandleAuthCallback(w http.ResponseWriter, r *http.Request) {
	token, err := c.service.GetUserToken(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	tokenCookie := &http.Cookie{Name: "accessToken", Value: *token, HttpOnly: false}
	http.SetCookie(w, tokenCookie)

	utils.SuccessMessageResponse(w, "Login Success")
}

func (c AuthController) HandleGetUserInfo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		fmt.Println("ERRORS: ", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	info, err := c.service.GetUserInfo(cookie.Value)
	if err != nil {
		fmt.Println("ERRORS: ", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	w.Write([]byte(info))
}

func (c AuthController) HandleLoginTemplate(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `
		<html>
			<body>
				<a href="/auth/login/google">Google Log In</a>
			</body>
		</html>`
	fmt.Fprintf(w, htmlIndex)
}