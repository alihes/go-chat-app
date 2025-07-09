package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Connect() error{
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://chatuser:1234@127.0.0.1:5432/chatapp?sslmodel=disable"
	}

	var err error
	Pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	return nil
}