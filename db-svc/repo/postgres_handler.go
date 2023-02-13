package repo

import (
	"context"
	"log"

	pgx "github.com/jackc/pgx/v5"
)

const (
	dockerConn = "postgres://postgres:maximilian@postgres_image:5432/maxDB?sslmode=disable"
	localConn  = "postgres://postgres:maximilian@localhost:5432/max?sslmode=disable"
)

func NewDatabaseConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dockerConn)
	if err != nil {
		log.Print(err)
	}
	return conn
}
