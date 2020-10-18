package pg_connection

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "senha"
	dbname   = "fialhoFabio_mc"
)

var connection *sql.DB

func Connection() *sql.DB {
	return connection
}

func Initialize() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	conn, err := sql.Open("postgres", connStr)
	connection = conn
	if err != nil {
		panic(err)
	}
}
