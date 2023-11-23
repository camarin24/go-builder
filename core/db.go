package core

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/mattn/go-sqlite3"
)

func NewDb() (*dbx.DB, error) {
	return dbx.Open("sqlite3", "./gostudio.sqlite3")
}