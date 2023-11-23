package database

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	*FactoryConfig
}

func NewPostgres(fc *FactoryConfig) *Postgres {
	return &Postgres{
		fc,
	}
}

func (p *Postgres) Ping() bool {
	db, err := dbx.Open("postgres", p.ConnStr)
	if err != nil {
		return false
	}
	err = db.DB().Ping()
	return err == nil
}
