package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lissteron/simplerr"
	"github.com/loghole/tracing"

	"github.com/lissteron/cloudsuny/internal/app/domain"
)

func (r *Repository) AddSession(ctx context.Context, session *domain.Session) (*domain.Session, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	session.CreatedAt = time.Now()
	session.UpdatedAt = time.Now()

	query := `INSERT INTO sessions (token, login, created_at, updated_at) VALUES (?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, session.Token, session.Login, session.CreatedAt, session.UpdatedAt)
	if err != nil {
		return nil, simplerr.Wrap(err, "can't insert")
	}

	return session, nil
}

func (r *Repository) FindSession(ctx context.Context, token string) (*domain.Session, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	query := `SELECT token, login, created_at, updated_at FROM sessions WHERE token = ?`

	dest := &domain.Session{}

	if err := r.db.GetContext(ctx, dest, query, token); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}

		return nil, simplerr.Wrap(err, "can't select")
	}

	return dest, nil
}

func (r *Repository) UpdateSession(ctx context.Context, session *domain.Session) error {
	defer tracing.ChildSpan(&ctx).Finish()

	session.UpdatedAt = time.Now()

	query := `UPDATE sessions SET updated_at = ? WHERE token = ?`

	if _, err := r.db.ExecContext(ctx, query, session.UpdatedAt, session.Token); err != nil {
		return simplerr.Wrap(err, "can't update")
	}

	return nil
}

func (r *Repository) DeleteSession(ctx context.Context, token string) error {
	defer tracing.ChildSpan(&ctx).Finish()

	query := `DELETE FROM sessions WHERE token = ?`

	if _, err := r.db.ExecContext(ctx, query, token); err != nil {
		return simplerr.Wrap(err, "can't update")
	}

	return nil
}
