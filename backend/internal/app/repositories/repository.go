package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var (
	ErrTxNotStarted    = errors.New("transaction not started")
	ErrInvalidDatabase = errors.New("invalid database")
)

type DB interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
}

type Repository struct {
	db     *sqlx.DB
	tx     *sqlx.Tx
	logger Logger
}

func NewRepository(db *sqlx.DB, logger Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func (r *Repository) database() DB {
	if r.tx != nil {
		return r.tx
	}

	return r.db
}

func (r *Repository) WithTransaction(ctx context.Context) (*Repository, error) {
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	return &Repository{db: r.db, tx: tx, logger: r.logger}, nil
}

func (r *Repository) CommitTransaction() error {
	if r.tx == nil {
		return ErrTxNotStarted
	}

	return r.tx.Commit()
}

func (r *Repository) RollbackTransaction() {
	if r.tx == nil {
		return
	}

	_ = r.tx.Rollback()
}
