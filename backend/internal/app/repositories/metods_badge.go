package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"

	"github.com/lissteron/cloudsuny/internal/app/domain"
	"github.com/lissteron/cloudsuny/internal/app/repositories/models"
)

func (r *Repository) CreateBadge(ctx context.Context, badge *domain.Badge) (*domain.Badge, error) {
	model := models.BadgeFromDomain(badge)
	model.ID = uuid.New().String()
	model.CreatedAt = time.Now()

	query := `INSERT INTO badge (id,user_id,type,x,y,created_at) VALUES (:id,:user_id,:type,:x,:y,:created_at)`

	if _, err := r.database().NamedExecContext(ctx, query, model); err != nil {
		return nil, err
	}

	return model.ToDomain(), nil
}

func (r *Repository) FindBadgeByID(ctx context.Context, badgeID string) (*domain.Badge, error) {
	query := "SELECT id,user_id,type,x,y,created_at FROM badge WHERE id = ?"

	var model = &models.Badge{}

	if err := r.database().GetContext(ctx, model, query, badgeID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return model.ToDomain(), nil
}

func (r *Repository) ListBadges(ctx context.Context) ([]*domain.Badge, error) {
	query := "SELECT id,user_id,type,x,y,created_at FROM badge"

	var dest []*models.Badge

	if err := r.database().SelectContext(ctx, &dest, query); err != nil {
		return nil, err
	}

	result := make([]*domain.Badge, len(dest))

	for idx, user := range dest {
		result[idx] = user.ToDomain()
	}

	return result, nil
}

func (r *Repository) UpdateBadge(ctx context.Context, badge *domain.Badge) (*domain.Badge, error) {
	model := models.BadgeFromDomain(badge)

	query := `UPDATE badge SET x=:x,y=:y WHERE id = :id`

	if _, err := r.database().NamedExecContext(ctx, query, model); err != nil {
		return nil, err
	}

	return model.ToDomain(), nil
}
