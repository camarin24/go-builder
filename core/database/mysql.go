package database

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
)

type MySql struct {
	*FactoryConfig
}

func NewMySql(fc *FactoryConfig) *MySql {
	return &MySql{fc}
}

func (my *MySql) Ping() bool {
	db, err := dbx.Open("mysql", my.ConnStr)
	if err != nil {
		return false
	}
	err = db.DB().Ping()
	return err == nil
}
