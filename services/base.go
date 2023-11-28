package services

import (
	"github.com/camarin24/go-studio/pkg/dbcontext"
	"github.com/camarin24/go-studio/pkg/log"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type Service struct {
	db  *dbcontext.DB
	log log.Logger
}

func NewService(db *dbx.DB, log log.Logger) *Service {
	return &Service{
		db:  dbcontext.New(db),
		log: log,
	}
}
