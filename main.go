package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gocraft/web"
)

//App main app context
type App struct {
}

//openidConfiguration (discovery)
func (a *App) openidConfiguration(rw web.ResponseWriter, r *web.Request) {

	data := map[string]interface{}{
		"issuer":                                           "https://" + Domain + "",
		"authorization_endpoint":                           "https://" + Domain + "/connect/authorize",
		"token_endpoint":                                   "https://" + Domain + "/connect/token",
		"token_endpoint_auth_signing_alg_values_supported": []string{"RS256"},
		"userinfo_endpoint":                                "https://" + Domain + "/connect/userinfo",
		"jwks_uri":                                         "https://" + Domain + "/jwks.json",
		"id_token_signing_alg_values_supported":            []string{"RS256"},
		"response_types_supported":                         []string{"code", "code id_token", "id_token"}}

	js, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(js))
}

//jwks (public rsa key)
func (a *App) jwks(rw web.ResponseWriter, r *web.Request) {

	data := map[string]string{
		"kty": "RSA",
		"alg": "RS256",
		"kid": string(time.Now().Unix()),
		"use": "sig",
		"e":   base64.StdEncoding.EncodeToString([]byte(string(RsaKey.PublicKey.E))),
		"n":   base64.StdEncoding.EncodeToString([]byte(RsaKey.PublicKey.N.String()))}

	js, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(js))
}

//RsaKey for jwks
var RsaKey *rsa.PrivateKey

//Domain .
var Domain string

func main() {

	RsaKey, _ = rsa.GenerateKey(rand.Reader, 256)

	Domain = "as.eko.ovh"
	app := App{}

	rootRouter := web.New(app)
	rootRouter.Middleware(web.LoggerMiddleware)
	rootRouter.Get(".well-known/openid-configuration", (*App).openidConfiguration)
	rootRouter.Get("jwks.json", (*App).jwks)
	rootRouter.Get("connect/authorize", (*App).openidConfiguration)
	rootRouter.Get("connect/token", (*App).openidConfiguration)
	rootRouter.Get("connect/userinfo", (*App).openidConfiguration)

	http.ListenAndServe("localhost:8080", rootRouter)

}
