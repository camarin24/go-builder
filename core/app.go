package core

import (
	"github.com/camarin24/go-studio/pkg/log"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type App struct {
	db  *dbx.DB
	log log.Logger
}

func NewApp() *App {
	db, err := NewDb()
	if err != nil {
		panic(err)
	}

	return &App{
		db:  db,
		log: log.New(),
	}
}

func (app *App) DB() *dbx.DB {
	return app.db
}

func (app *App) Log() log.Logger {
	return app.log
}
