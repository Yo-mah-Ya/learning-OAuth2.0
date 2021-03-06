package main

import "html/template"

const (
	//SCOPE                 = "readonly"
	SCOPE                 = "https://www.googleapis.com/auth/photoslibrary.readonly"
	AUTH_CODE_DURATION    = 300
	ACCESS_TOKEN_DURATION = 3600
)

type Client struct {
	id          string
	name        string
	redirectURL string
	secret      string
}

type User struct {
	id          int
	name        string
	password    string
	sub         string
	given_name  string
	family_name string
	locale      string
}

type Session struct {
	client                string
	state                 string
	scopes                string
	redirectUri           string
	code_challenge        string
	code_challenge_method string
	// OIDC
	nonce string
	// whether to publish ID token. if false publish just only OAuth token
	oidc bool
}

type AuthCode struct {
	user         string
	clientId     string
	scopes       string
	redirect_uri string
	expires_at   int64
}

type TokenCode struct {
	user       string
	clientId   string
	scopes     string
	expires_at int64
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	IdToken     string `json:"id_token,omitempty"`
}

type Payload struct {
	Iss        string `json:"iss"`
	Azp        string `json:"azp"`
	Aud        string `json:"aud"`
	Sub        string `json:"sub"`
	AtHash     string `json:"at_hash"`
	Nonce      string `json:"nonce"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Locale     string `json:"locale"`
	Iat        int64  `json:"iat"`
	Exp        int64  `json:"exp"`
}

type jwkKey struct {
	Kid string `json:"kid"`
	N   string `json:"n"`
	Alg string `json:"alg"`
	Kty string `json:"kty"`
	E   string `json:"e"`
	Use string `json:"use"`
}

var templates = make(map[string]*template.Template)
var sessionList = make(map[string]Session)
var AuthCodeList = make(map[string]AuthCode)
var TokenCodeList = make(map[string]TokenCode)

var clientInfo = Client{
	id:          "1234",
	name:        "test",
	redirectURL: "http://localhost:8080/callback",
	secret:      "secret",
}

var user = User{
	id:          1111,
	name:        "Christoph Jenkins",
	password:    "password",
	sub:         "11111111",
	given_name:  "Christoph",
	family_name: "Jenkins",
	locale:      "New York",
}
