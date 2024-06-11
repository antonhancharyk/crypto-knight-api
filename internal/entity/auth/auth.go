package auth

import "crypto/rsa"

type QueryParams struct {
	Token string `form:"token" binding:"required"`
}

type RSAKeys struct {
	PublicKey *rsa.PublicKey
}
