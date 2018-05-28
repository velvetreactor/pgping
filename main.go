package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Connection interface {
	Open(string, string) (*sql.DB, error)
}

type Logger interface {
	Print(...interface{})
}

type DBPinger interface {
	Ping() error
}

type SystemLogger struct{}

func (sl *SystemLogger) Print(v ...interface{}) {
	log.Print(v)
}

type PostgresConnection struct{}

func (pc *PostgresConnection) Open(driverName, dataSourceName string) (*sql.DB, error) {
	return sql.Open("postgres", dataSourceName)
}

type SystemDBPinger struct {
	db *sql.DB
}

func (sdbp *SystemDBPinger) Ping() error {
	return sdbp.db.Ping()
}

func main() {
	// err = db.Ping()
	// if err != nil {
	// 	log.Print(err)
	// 	os.Exit(1)
	// }
	// log.Print("Postgres server READY and accepting connections...")
	// os.Exit(0)
	pc := &PostgresConnection{}
	os.Exit(
		realMain(
			pc,
			&SystemLogger{},
			&DBPing{},
		),
	)
}

func realMain(c Connection, l Logger, dp DBPinger) int {
	cs := os.Getenv("PGCONN")
	if cs == "" {
		l.Print("PGCONN environment variable cannot be blank")
		return 1
	}
	con, err := c.Open("postgres", cs)
	if err != nil {
		l.Print(err)
		return 1
	}
	return 0
}
