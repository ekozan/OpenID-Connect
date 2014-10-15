package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	// jwt "github.com/dgrijalva/jwt-go"
	"github.com/gocraft/web"
)

//App main app context
type App struct {
}

//Auth .
func (a *App) Auth(rw web.ResponseWriter, r *web.Request, next web.NextMiddlewareFunc) {
	// rw.WriteHeader(http.StatusUnauthorized)
	// token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
	// 	return myLookupKey(token.Header["kid"])
	// })
	// token := jwt.New(jwt.GetSigningMethod("HS512"))
	// // Set some claims
	// token.Claims["iss"] = "auth.eko.ovh"
	// token.Claims["kid"] = "jioedjaziodjziodjefuojfi"
	// token.Claims["access"] = "official"
	// token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// // Sign and get the complete encoded token as a string
	// tokenString, _ := token.SignedString([]byte("SecretKey"))
	// fmt.Fprint(rw, tokenString)

	next(rw, r)
}

//openidConfiguration (discovery)
func (a *App) openidConfiguration(rw web.ResponseWriter, r *web.Request) {

	data := map[string]string{
		"issuer":                 "https://" + Domain + "",
		"authorization_endpoint": "https://" + Domain + "/connect/authorize",
		"token_endpoint":         "https://" + Domain + "/connect/token",
		"userinfo_endpoint":      "https://" + Domain + "/connect/userinfo",
		"jwks_uri":               "https://" + Domain + "/jwks.json"}

	js, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(js))
}

//jwks .
func (a *App) jwks(rw web.ResponseWriter, r *web.Request) {

	data := map[string]string{
		"kty": "RSA",
		"alg": "RS256",
		"use": "sig",
		"e":   base64.StdEncoding.EncodeToString([]byte(string(RsaKey.PublicKey.E))),
		"n":   base64.StdEncoding.EncodeToString([]byte(RsaKey.PublicKey.N.String()))}

	js, _ := json.Marshal(data)
	fmt.Fprint(rw, string(js))
}

// func (a *App) initDb() {
//
// 	db, err := sql.Open("sqlite3", "db.bin")
// 	checkErr(err, "sql.Open failed")
// 	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
//
// 	dbmap.AddTableWithName(Card{}, "cards").SetKeys(false, "UUID")
// 	dbmap.AddTableWithName(Skill{}, "skills").SetKeys(false, "UUID")
//
// 	err = dbmap.CreateTablesIfNotExists()
// 	checkErr(err, "Create tables failed")
//
// 	a.dbmap = dbmap
// }

//RsaKey for jwks

var RsaKey *rsa.PrivateKey
var Domain string

func main() {

	RsaKey, _ = rsa.GenerateKey(rand.Reader, 256)

	Domain = "as.eko.ovh"
	app := App{}

	rootRouter := web.New(app)
	rootRouter.Middleware((*App).Auth).Middleware(web.LoggerMiddleware)
	rootRouter.Get(".well-known/openid-configuration", (*App).openidConfiguration)
	rootRouter.Get("jwks.json", (*App).jwks)

	http.ListenAndServe("localhost:8080", rootRouter)

}
