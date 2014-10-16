package OpenIDConnect

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gocraft/web"
)

var (
	ErrInvalidKey      = errors.New("key is invalid or of invalid type")
	ErrHashUnavailable = errors.New("the requested hash function is unavailable")
)

//New initialise App context
func New(s Storage) *App {
	return &App{Storage: s, Config: s.getConfig()}
}

//App main app context
type App struct {
	Storage Storage
	Config  Config
}

//Config of App
type Config struct {
	domain string
}

//Client -> oidApplications
type Client struct {
	Name string
}

//AccessToken -> oid token
type AccessToken struct {
	Client       string
	User         string
	Token        string
	RefreshToken string //Optional
	Expire       string //Optional
	CreatedAt    string
}

//jwks (public rsa key)
func (a *App) jwks(rw web.ResponseWriter, r *web.Request) {

	data := map[string]string{
		"kty": "RSA",
		"alg": "RS256",
		"kid": string(time.Now().Unix()),
		"use": "sig"}
	// "e":   base64.StdEncoding.EncodeToString([]byte(string(RsaKey.PublicKey.E))),
	// "n":   base64.StdEncoding.EncodeToString([]byte(RsaKey.PublicKey.N.String()))}

	js, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(js))
}

/////JWT

//Jwt struct
type Jwt struct {
	Raw    string
	Header map[string]string
	Claims map[string]string
	Sign   string
	Valid  bool
}

func parseJwt(jwts string) (*Jwt, error) {
	parts := strings.Split(jwts, ".")
	if len(parts) != 3 {
		return nil, ErrInvalidKey
	}

	token := &Jwt{Raw: jwts}

	return token, nil
}

func (jwt *Jwt) verify() bool {
	return false
}

func (jwt *Jwt) encode() {
}

type jwtSignEngine interface {
	verify(token string) bool
	sign() string
	algo() string
}
