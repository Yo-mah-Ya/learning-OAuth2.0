package token

import (
	"api/openapi"
	"api/shared"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type TokenServer struct{}

type Response struct {
	status int
	body   []byte
}

func service(params openapi.TokenParams) Response {
	v, ok := shared.AuthCodeList[params.Code]
	if !ok {
		log.Println("auth code isn't exist")
		return Response{
			status: http.StatusBadRequest,
			body:   []byte(fmt.Sprintf("no authorization code")),
		}
	}

	delete(shared.AuthCodeList, params.Code)

	tokenResp := struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int64  `json:"expires_in"`
		IdToken     string `json:"id_token,omitempty"`
	}{
		AccessToken: uuid.New().String(),
		TokenType:   "Bearer",
		ExpiresIn:   time.Now().Unix() + shared.ACCESS_TOKEN_DURATION,
	}

	if shared.SessionList[*params.Session].Oidc {
		tokenResp.IdToken, _ = makeJWT()
	}
	resp, err := json.Marshal(tokenResp)
	if err != nil {
		log.Println("json marshal err")
		return Response{
			status: http.StatusInternalServerError,
			body:   []byte(http.StatusText(http.StatusInternalServerError)),
		}
	}

	log.Printf("token ok to client %s, token is %s", v.ClientId, string(resp))
	return Response{
		status: http.StatusOK,
		body:   resp,
	}
}

func (s *TokenServer) Token(w http.ResponseWriter, r *http.Request, params openapi.TokenParams) {
	response := service(params)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.status)
	w.Write(response.body)
}
