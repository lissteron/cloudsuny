package repositories

import (
	"github.com/loghole/database"
	"github.com/loghole/tracing/tracelog"
)

type Repository struct {
	db     *database.DB
	logger tracelog.Logger
}

func NewRepository(db *database.DB, logger tracelog.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}
