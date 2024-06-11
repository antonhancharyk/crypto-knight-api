package auth

import (
	"github.com/antongoncharik/crypto-knight-api/internal/entity/auth"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
	"github.com/antongoncharik/crypto-knight-api/pkg/utilities"
)

type Auth struct {
	repo *repository.Repository
	keys auth.RSAKeys
}

func New(repo *repository.Repository, keys auth.RSAKeys) *Auth {
	return &Auth{repo, keys}
}

func (a *Auth) ValidateToken(token string) error {
	return utilities.ValidateToken(token, a.keys.PublicKey)

}
