package connection

import (
	"context"
	"fmt"

	"github.com/gabehamasaki/encurtago/internal/config"
	"github.com/jackc/pgx/v5"
)

func NewConnection(ctx context.Context, cfg *config.Config) (*pgx.Conn, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s host=%s password=%s sslmode=disable",
		cfg.DB_USER, cfg.DB_NAME, cfg.DB_HOST, cfg.DB_PASSWORD)

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
