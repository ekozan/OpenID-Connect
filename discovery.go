package OpenIDConnect

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//openidConfiguration (discovery)
func (a *App) openidConfiguration(rw http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"issuer":                                           "https://" + a.Config.domain + "",
		"authorization_endpoint":                           "https://" + a.Config.domain + "/connect/authorize",
		"token_endpoint":                                   "https://" + a.Config.domain + "/connect/token",
		"token_endpoint_auth_signing_alg_values_supported": []string{"RS256"},
		"userinfo_endpoint":                                "https://" + a.Config.domain + "/connect/userinfo",
		"jwks_uri":                                         "https://" + a.Config.domain + "/jwks.json",
		"id_token_signing_alg_values_supported":            []string{"RS256"},
		"response_types_supported":                         []string{"code", "code id_token", "id_token"}}

	js, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(js))
}
