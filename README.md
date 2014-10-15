#Golang OpenID Connect

Golang implementation of OpenID Connect 1.0 (http://openid.net/connect/)

### suported response type (configurable by client)
* `code` 	  Authorization Code Flow
* `id_token` 	Implicit Flow
* `id_token token` 	Implicit Flow
* `code id_token` 	Hybrid Flow
* `code token` 	Hybrid Flow
* `code id_token token` 	Hybrid Flow

### suported token
`Token` `Token_id` `Refresh_token`

##RoadMap

v0.0.1 : (october 14)
* Minimal implementation (core)

v0.0.2 : (october 14)
* Start of Dynamic implementation (Discovery)

v0.0.3 : (november 14)
* Full support of dynamic

v1.0.0 : (december 14)
* code clean
* finish all test


##Installation
`go get github.com/3ko/openidconnect`

##Configuration
