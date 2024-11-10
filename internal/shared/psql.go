package shared

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"strconv"
)

const pgDriverName = "postgres"

func NewPSQL() *sql.DB {

	db, err := sql.Open(
		pgDriverName,
		buildDataSourceName(
			"localhost",
			"go_whisky",
			"public",
			"go_whisky",
			"",
			"disable",
		),
	)
	if err != nil {
		panic(err)
	}
	maxConnections, err := strconv.Atoi("10")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(maxConnections)
	db.SetMaxIdleConns(maxConnections)

	if err != nil {
		panic(err)
	}

	boil.DebugMode = true
	boil.SetDB(db)
	return db
}

func buildDataSourceName(host string, dbName string, schemaName string, user string, password string, sslMode string) string {
	return fmt.Sprintf("host=%s dbname=%s search_path=%s user=%s password='%s' sslmode=%s", host, dbName, schemaName, user, password, sslMode)
}
