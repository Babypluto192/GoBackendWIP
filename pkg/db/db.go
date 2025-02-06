package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"sync"
)

type DB struct {
	mu   sync.Mutex
	Pool *pgxpool.Pool
}

func New(connectionString string) (*DB, error) {
	slog.Debug("Подключаюсь к", connectionString)
	slog.Info("Подключась к БД")
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info("Успешно подключился к БД")
	return &DB{Pool: pool}, nil
}
