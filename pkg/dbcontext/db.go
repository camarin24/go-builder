package dbcontext

// This file is based on https://github.com/qiangxue/go-rest-api/blob/master/pkg/dbcontext/db.go

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

const (
	txKey int = iota
)

type DB struct {
	db *dbx.DB
}

func New(db *dbx.DB) *DB {
	return &DB{db}
}

func (db *DB) DB() *dbx.DB {
	return db.db
}

func (db *DB) With(ctx context.Context) dbx.Builder {
	if tx, ok := ctx.Value(txKey).(*dbx.Tx); ok {
		return tx
	}
	return db.db.WithContext(ctx)
}
