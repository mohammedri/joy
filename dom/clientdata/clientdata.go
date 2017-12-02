package clientdata

import "github.com/matthewmueller/golly/dom/webauthnextensions"

type ClientData struct {
	challenge    string
	extensions   *webauthnextensions.WebAuthnExtensions
	hashAlg      interface{}
	origin       string
	rpId         string
	tokenBinding *string
}