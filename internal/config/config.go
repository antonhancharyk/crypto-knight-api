package config

import (
	"os"

	"github.com/antongoncharik/crypto-knight-api/internal/entity/auth"
	"github.com/antongoncharik/crypto-knight-api/pkg/utilities"
)

type DBConfig struct {
	User     string
	Name     string
	Password string
	Host     string
	Port     string
}

type GRPCConfig struct {
	Host string
}

type HTTPConfig struct {
	Addr string
}

type Config struct {
	Env   string
	Auth  auth.RSAKeys
	DB    DBConfig
	GRPC  GRPCConfig
	HTTP  HTTPConfig
}

func MustLoad() (Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	publicKey, err := utilities.LoadPublicKey("./config/rsa/public_key.pem")
	if err != nil {
		return Config{}, err
	}

	dbCfg := DBConfig{
		User:     os.Getenv("DB_USER"),
		Name:     os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}

	grpcCfg := GRPCConfig{
		Host: os.Getenv("GRPC_HOST"),
	}

	httpAddr := os.Getenv("HTTP_ADDR")
	if httpAddr == "" {
		httpAddr = ":8080"
	}

	httpCfg := HTTPConfig{
		Addr: httpAddr,
	}

	return Config{
		Env:  env,
		Auth: auth.RSAKeys{PublicKey: publicKey},
		DB:   dbCfg,
		GRPC: grpcCfg,
		HTTP: httpCfg,
	}, nil
}

