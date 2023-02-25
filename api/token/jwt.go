package token

import (
	"api/shared"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/lestrrat-go/jwx/jwk"
)

func readPrivateKey() (*rsa.PrivateKey, error) {
	data, err := ioutil.ReadFile("private-key.pem")
	if err != nil {
		return nil, err
	}
	keyBlock, _ := pem.Decode(data)
	if keyBlock == nil {
		return nil, fmt.Errorf("invalid private key data")
	}
	if keyBlock.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("invalid private key type : %s", keyBlock.Type)
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func makeHeaderPayload() string {
	payload_json, _ := json.Marshal(struct {
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
	}{
		Iss:        "https://oreore.oidc.com",
		Azp:        shared.ClientInfo.Id,
		Aud:        shared.ClientInfo.Id,
		Sub:        shared.TestUser.Sub,
		AtHash:     "PRzSZsEPQVqzY8xyB2ls5A",
		Nonce:      "abc",
		Name:       shared.TestUser.Name,
		GivenName:  shared.TestUser.Given_name,
		FamilyName: shared.TestUser.Family_name,
		Locale:     shared.TestUser.Locale,
		Iat:        time.Now().Unix(),
		Exp:        time.Now().Unix() + shared.ACCESS_TOKEN_DURATION,
	})

	return fmt.Sprintf(
		"%s.%s",
		base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"RS256","kid": "12345678","typ":"JWT"}`)),
		base64.RawURLEncoding.EncodeToString(payload_json),
	)
}

func makeJWT() (string, error) {
	jwtString := makeHeaderPayload()

	privateKey, err := readPrivateKey()
	if err != nil {
		return "", err
	}
	if err := privateKey.Validate(); err != nil {
		return "", fmt.Errorf("private key validate err : %s", err)
	}
	hash := sha256.New()
	hash.Write([]byte(jwtString))
	tokenHash := hash.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, tokenHash)
	if err != nil {
		return "", fmt.Errorf("sign by private key is err : %s", err)
	}

	return fmt.Sprintf("%s.%s", jwtString, base64.RawURLEncoding.EncodeToString(signature)), nil
}

func makeJWK() []byte {

	data, _ := ioutil.ReadFile("public-key.pem")
	keyset, _ := jwk.ParseKey(data, jwk.WithPEM(true))

	keyset.Set(jwk.KeyIDKey, "12345678")
	keyset.Set(jwk.AlgorithmKey, "RS256")
	keyset.Set(jwk.KeyUsageKey, "sig")

	buf, _ := json.MarshalIndent(map[string]interface{}{
		"keys": []interface{}{keyset},
	}, "", "  ")
	return buf

}
