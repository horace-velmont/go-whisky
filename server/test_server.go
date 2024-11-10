package server

import (
	"context"
	"database/sql"
	"github.com/GagulProject/go-whisky/internal/repositoryfx"
	"github.com/GagulProject/go-whisky/internal/servicefx"
	"github.com/GagulProject/go-whisky/internal/sharedfx"
	"go.uber.org/fx"
)

type TestServer struct {
	app *fx.App
	db  *sql.DB
}

var Populate = fx.Populate

func NewTestServer(testOpts ...fx.Option) *TestServer {
	var db *sql.DB

	fxApp := fx.New(
		append(
			testOpts,
			TestOptions(),
			Populate(&db),
		)...,
	)

	return &TestServer{
		app: fxApp,
		db:  db,
	}
}

func (t *TestServer) WithPreparedTables(tableNames ...string) *TestServer {
	for _, tableName := range tableNames {
		_, err := t.db.Exec("TRUNCATE TABLE " + tableName + " CASCADE")
		if err != nil {
			panic(err)
		}
	}
	return t
}

func (t *TestServer) Stop() {
	err := t.app.Stop(context.Background())
	if err != nil {
		return
	}
}

func TestOptions() fx.Option {
	return fx.Options(
		servicefx.Option,
		repositoryfx.Option,
		sharedfx.Option,
	)
}
