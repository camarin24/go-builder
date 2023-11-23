package core

import dbx "github.com/go-ozzo/ozzo-dbx"

type BaseApp struct {
	db *dbx.DB
}

func NewBaseApp() *BaseApp {
	return &BaseApp{}
}

func (ba *BaseApp) DB() *dbx.DB {
	if ba.db == nil {
		db, err := NewDb()
		if err != nil {
			panic(err)
		}
		ba.db = db
	}
	return ba.db
}
