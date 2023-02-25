package auth

import (
	"api/openapi"
	"api/shared"
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
)

const html = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Login</title>
</head>
<body>
    <h2>Client App</h2>
    <p>{{.ClientId}}</p>
    <h2>Requested Permissions</h2>
    <p>{{.Scope}}</p>
    <hr>
    <div align="center">
        <table border="0">
            <form action="/authcheck" method="post">
                <tr>
                    <th>
                        user id
                    </th>
                    <td>
                        <input type="text" name="username" value="" size="24">
                    </td>
                </tr>
                <tr>
                    <th>
                        password
                    </th>
                    <td>
                        <input type="password" name="password" value="" size="24">
                    </td>
                </tr>
                <tr>
                    <td colspan="2">
                        <input type="submit" value="allow">
                    </td>
                </tr>
            </form>
        </table>
    </div>
</body>
</html>
`

type AuthServer struct{}

func (s *AuthServer) Auth(w http.ResponseWriter, r *http.Request, params openapi.AuthParams) {
	session := shared.Session{
		Client:      params.ClientId,
		State:       params.State,
		Scopes:      params.Scope,
		RedirectUri: params.RedirectUri,
	}

	sessionId := uuid.New().String()
	shared.SessionList[sessionId] = session

	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: sessionId,
	})

	if err := template.Must(template.New("login").Parse(html)).Execute(w, struct {
		ClientId string
		Scope    string
	}{
		ClientId: session.Client,
		Scope:    session.Scopes,
	}); err != nil {
		log.Println(err)
	}
}
