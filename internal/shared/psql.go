package shared

import (
	"database/sql"
	"fmt"
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

// password 없이 사용하는 경우 escape가 없으면 sslmode 에러가 발생함
func buildDataSourceName(host string, dbName string, schemaName string, user string, password string, sslMode string) string {
	return fmt.Sprintf("host=%s dbname=%s search_path=%s user=%s password='%s' sslmode=%s", host, dbName, schemaName, user, password, sslMode)
}
