package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/lissteron/cloudsuny/internal/app/domain"
	"github.com/lissteron/cloudsuny/internal/app/repositories/models"
)

func (r *Repository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	model := models.UserFromDomain(user)
	model.ID = uuid.New().String()
	model.CreatedAt = time.Now()

	query := `INSERT INTO user (id,username,avatar,created_at) VALUES (:id,:username,:avatar,:created_at)`

	if _, err := r.database().NamedExecContext(ctx, query, model); err != nil {
		return nil, err
	}

	return model.ToDomain(), nil
}

func (r *Repository) FindUserByID(ctx context.Context, userID string) (*domain.User, error) {
	query := "SELECT id,username,avatar,created_at FROM user WHERE id = ?"

	var model = &models.User{}

	if err := r.database().GetContext(ctx, model, query, userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return model.ToDomain(), nil
}

func (r *Repository) FindUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	query := "SELECT id,username,avatar,created_at FROM user WHERE username = ?"

	var model = &models.User{}

	if err := r.database().GetContext(ctx, model, query, username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return model.ToDomain(), nil
}

func (r *Repository) ListUsers(ctx context.Context) ([]*domain.User, error) {
	query := "SELECT id,username,avatar,created_at FROM user"

	var dest []*models.User

	if err := r.database().SelectContext(ctx, &dest, query); err != nil {
		return nil, err
	}

	result := make([]*domain.User, len(dest))

	for idx, user := range dest {
		result[idx] = user.ToDomain()
	}

	return result, nil
}
