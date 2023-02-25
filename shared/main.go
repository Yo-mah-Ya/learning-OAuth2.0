package shared

const (
	ACCESS_TOKEN_DURATION = 3600
)

type Client struct {
	Id string
}

type Session struct {
	Client                string
	State                 string
	Scopes                string
	RedirectUri           string
	Code_challenge        string
	Code_challenge_method string
	// OIDC
	Nonce string
	// whether to publish ID token. if false publish just only OAuth token
	Oidc bool
}

type AuthCode struct {
	User         string
	ClientId     string
	Scopes       string
	Redirect_uri string
	Expires_at   int64
}

var SessionList = make(map[string]Session)
var AuthCodeList = make(map[string]AuthCode)

var ClientInfo = Client{
	Id: "1234",
}

type User struct {
	Id          int
	Name        string
	Password    string
	Sub         string
	Given_name  string
	Family_name string
	Locale      string
}

var TestUser = User{
	Id:          1111,
	Name:        "u",
	Password:    "p",
	Sub:         "11111111",
	Given_name:  "Christoph",
	Family_name: "Jenkins",
	Locale:      "New York",
}
