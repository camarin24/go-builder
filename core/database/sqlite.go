package database

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	*FactoryConfig
}

func NewSqlite(fc *FactoryConfig) *Sqlite {
	return &Sqlite{fc}
}

func (sq *Sqlite) Ping() bool {
	db, err := dbx.Open("sqlite3", sq.ConnStr)
	if err != nil {
		return false
	}
	err = db.DB().Ping()
	return err == nil
}
