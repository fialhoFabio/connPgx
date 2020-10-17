package pg_connection

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "senha"
	dbname   = "fialhoFabio_mc"
)

var connection *pgxpool.Pool

func Connection() *pgxpool.Pool {
	return connection
}

func Initialize() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	conn, err := pgxpool.Connect(context.Background(), connStr)
	connection = conn
	if err != nil {
		panic(err)
	}
}
