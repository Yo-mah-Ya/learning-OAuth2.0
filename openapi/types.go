// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package openapi

// Defines values for AuthParamsResponseType.
const (
	Code AuthParamsResponseType = "code"
)

// Message defines model for Message.
type Message struct {
	Message string `json:"message"`
}

// AuthParams defines parameters for Auth.
type AuthParams struct {
	// ClientId The ID for the desired user pool app client.
	ClientId string `form:"client_id" json:"client_id"`

	// State (optional but recommended) - A random value that's used to prevent cross-site request forgery (CSRF) attacks.
	State string `form:"state" json:"state"`

	// Scope A space-separated list of scopes to request for the generated tokens. Note that:
	//   - An ID token is only generated if the openid scope is requested.
	//   - The phone, email, and profile scopes can only be requested if openid is also requested.
	//   - A vended access token can only be used to make user pool API calls if aws.cognito.signin.user.admin is requested.
	Scope string `form:"scope" json:"scope"`

	// RedirectUri The URL that a user is directed to after successful authentication.
	RedirectUri string `form:"redirect_uri" json:"redirect_uri"`

	// ResponseType Set to “code” for this grant type.
	ResponseType AuthParamsResponseType `form:"response_type" json:"response_type"`
}

// AuthParamsResponseType defines parameters for Auth.
type AuthParamsResponseType string

// AuthCheckFormdataBody defines parameters for AuthCheck.
type AuthCheckFormdataBody struct {
	// Password password
	Password string `json:"password"`

	// Username username
	Username string `json:"username"`
}

// TokenParams defines parameters for Token.
type TokenParams struct {
	// GrantType Set to “authorization_code” for this grant.
	GrantType string `form:"grant_type" json:"grant_type"`

	// Code The authorization code that's vended to the user.
	Code string `form:"code" json:"code"`

	// ClientId Same as from the request in step 1.
	ClientId string `form:"client_id" json:"client_id"`

	// RedirectUri Same as from the request in step 1.
	RedirectUri string `form:"redirect_uri" json:"redirect_uri"`

	// ClientSecret (optional, is required if a code_challenge was specified in the original request) – The base64 URL-encoded representation of the unhashed, random string that was used to generate the PKCE code_challenge in the original request.
	ClientSecret string `form:"client_secret" json:"client_secret"`

	// Session session
	Session *string `form:"session,omitempty" json:"session,omitempty"`
}

// AuthCheckFormdataRequestBody defines body for AuthCheck for application/x-www-form-urlencoded ContentType.
type AuthCheckFormdataRequestBody AuthCheckFormdataBody
