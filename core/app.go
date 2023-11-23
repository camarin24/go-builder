package core

import dbx "github.com/go-ozzo/ozzo-dbx"

type App interface {
	DB() *dbx.DB
}
