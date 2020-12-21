package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lissteron/simplerr"
	"github.com/loghole/tracing"

	"github.com/lissteron/cloudsuny/internal/app/domain"
	"github.com/lissteron/cloudsuny/internal/app/repositories/models"
)

func (r *Repository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	model := models.UserFromDomain(user)
	model.ID = uuid.New().String()
	model.CreatedAt = time.Now()

	var (
		query  = `SELECT EXISTS (SELECT 1 FROM user WHERE username = ?)`
		exists bool
	)

	if err := r.db.GetContext(ctx, &exists, query, user.Username); err != nil {
		return nil, simplerr.Wrap(err, "can't check exists")
	}

	if exists {
		return nil, domain.ErrAlreadyExists
	}

	query = `INSERT INTO user (id,username,avatar,created_at) VALUES (:id,:username,:avatar,:created_at)`

	if _, err := r.db.NamedExecContext(ctx, query, model); err != nil {
		return nil, simplerr.Wrap(err, "can't insert")
	}

	return model.ToDomain(), nil
}

func (r *Repository) FindUserByID(ctx context.Context, userID string) (*domain.User, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	query := "SELECT id,username,avatar,created_at FROM user WHERE id = ?"

	model := &models.User{}

	if err := r.db.GetContext(ctx, model, query, userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}

		return nil, simplerr.Wrap(err, "can't select")
	}

	return model.ToDomain(), nil
}

func (r *Repository) ListUsers(ctx context.Context, limit, offset uint64) (*domain.UserListResult, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	query := "SELECT id,username,avatar,created_at FROM user LIMIT ? OFFSET ?"

	var dest []*models.User

	if err := r.db.SelectContext(ctx, &dest, query, limit, offset); err != nil {
		return nil, simplerr.Wrap(err, "can't select")
	}

	result := &domain.UserListResult{
		List:    make([]*domain.User, len(dest)),
		Count:   int64(len(dest)),
		Skipped: int64(offset),
	}

	for idx, user := range dest {
		result.List[idx] = user.ToDomain()
	}

	query = "SELECT count(id) FROM user"

	if err := r.db.GetContext(ctx, &result.Total, query); err != nil {
		return nil, simplerr.Wrap(err, "can't get count")
	}

	return result, nil
}

func (r *Repository) DeleteUser(ctx context.Context, userID string) error {
	defer tracing.ChildSpan(&ctx).Finish()

	err := r.db.RunTxx(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		var (
			query  = "SELECT EXISTS (SELECT 1 FROM user WHERE id = ?)"
			exists bool
		)

		if err := tx.GetContext(ctx, &exists, query, userID); err != nil {
			return simplerr.Wrap(err, "select user")
		}

		if !exists {
			return domain.ErrNotFound
		}

		query = "DELETE FROM badge WHERE user_id = ?"

		if _, err := tx.ExecContext(ctx, query, userID); err != nil {
			return simplerr.Wrap(err, "delete badges")
		}

		query = "DELETE FROM user WHERE id = ?"

		if _, err := tx.ExecContext(ctx, query, userID); err != nil {
			return simplerr.Wrap(err, "delete user")
		}

		return nil
	})
	if err != nil {
		return simplerr.Wrap(err, "transaction err")
	}

	return nil
}
