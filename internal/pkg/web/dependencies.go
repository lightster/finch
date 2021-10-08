package web

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Dependencies struct {
	Config *Config
}

func NewDependencies(config *Config) *Dependencies {
	return &Dependencies{Config: config}
}

func (d *Dependencies) InitDB(ctx context.Context) (*pgx.Conn, error) {
	db, err := pgx.Connect(ctx, d.Config.DatabaseURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
