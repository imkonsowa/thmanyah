package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"thmanyah/internal/conf"
)

func NewPgPool(ctx context.Context, conf *conf.Data) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		conf.Postgres.Host,
		conf.Postgres.User,
		conf.Postgres.Password,
		conf.Postgres.Dbname,
		conf.Postgres.Port,
	)

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
