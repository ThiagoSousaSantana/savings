package db

import (
	"context"
	"database/sql"

	"github.com/ThiagoSousaSantana/saving/cmd/config"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type DB struct {
	conn *sql.DB
}

func NewDB(lc fx.Lifecycle, log *zap.Logger, config *config.Config) (*DB, error) {
	var conn *sql.DB

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			conn, err := sql.Open("postgres", "host="+config.DB.Host+" port="+config.DB.Port+" user="+config.DB.User+" password="+config.DB.Password+" dbname="+config.DB.Database+" sslmode=disable")
			if err != nil {
				log.Error("Error connecting to database", zap.Error(err))
				return err
			}

			err = conn.Ping()
			if err != nil {
				log.Error("Error pinging database", zap.Error(err))
				return err
			}

			log.Info("Database connected")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := conn.Close()
			if err != nil {
				log.Error("Error closing database connection", zap.Error(err))
				return err
			}

			log.Info("Database connection closed")
			return nil
		},
	})

	return &DB{
		conn: conn,
	}, nil
}
