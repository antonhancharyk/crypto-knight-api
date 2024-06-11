package config

import (
	"github.com/antongoncharik/crypto-knight-api/internal/entity/auth"
	"github.com/antongoncharik/crypto-knight-api/pkg/utilities"
)

func MustLoad() (auth.RSAKeys, error) {
	publicKey, err := utilities.LoadPublicKey("./config/rsa/public_key.pem")
	if err != nil {
		return auth.RSAKeys{}, err
	}

	return auth.RSAKeys{PublicKey: publicKey}, nil
}
