package repo

import (
	"context"
	"log"
	"max-db-svc/config"

	pgx "github.com/jackc/pgx/v5"
)

func NewDatabaseConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), config.Conf.Postgres.ConnectionString)
	if err != nil {
		log.Print(err)
	}
	return conn
}
