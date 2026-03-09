package database

import (
	"fmt"

	"github.com/antongoncharik/crypto-knight-api/internal/config"
	"github.com/antongoncharik/crypto-knight-api/internal/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(cfg config.DBConfig) (*sqlx.DB, error) {
	connectStr := fmt.Sprintf(
		"user=%s dbname=%s sslmode=disable password=%s host=%s port=%s",
		cfg.User,
		cfg.Name,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)

	db, err := sqlx.Connect("postgres", connectStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	logger.Log.Infow("DB is running")

	return db, nil
}
