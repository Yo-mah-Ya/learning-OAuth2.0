package authCheck

import (
	"api/openapi"
	"api/shared"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type AuthCheckServer struct{}

func (s *AuthCheckServer) AuthCheck(w http.ResponseWriter, r *http.Request) {
	userInfo := openapi.AuthCheckFormdataBody{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	if userInfo.Username != shared.TestUser.Name || userInfo.Password != shared.TestUser.Password {
		fmt.Println("Wrong User Info!!")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(openapi.Message{
			Message: "Bad Request",
		})
	}
	cookie, _ := r.Cookie("session")
	http.SetCookie(w, cookie)
	v, _ := shared.SessionList[cookie.Value]

	authCodeString := uuid.New().String()

	shared.AuthCodeList[authCodeString] = shared.AuthCode{
		User: openapi.AuthCheckFormdataBody{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}.Username,
		ClientId:     v.Client,
		Scopes:       v.Scopes,
		Redirect_uri: v.RedirectUri,
		Expires_at:   time.Now().Unix() + 300,
	}

	w.Header().Add("Location", fmt.Sprintf("%s?code=%s&state=%s", v.RedirectUri, authCodeString, v.State))
	w.WriteHeader(302)
}
