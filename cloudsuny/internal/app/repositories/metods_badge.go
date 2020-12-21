package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lissteron/simplerr"
	"github.com/loghole/tracing"

	"github.com/lissteron/cloudsuny/internal/app/domain"
	"github.com/lissteron/cloudsuny/internal/app/repositories/models"
)

func (r *Repository) CreateBadge(ctx context.Context, badge *domain.Badge) (*domain.Badge, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	model := models.BadgeFromDomain(badge)
	model.ID = uuid.New().String()
	model.CreatedAt = time.Now()

	query := `INSERT INTO badge (id,user_id,type,x,y,created_at) VALUES (:id,:user_id,:type,:x,:y,:created_at)`

	if _, err := r.db.NamedExecContext(ctx, query, model); err != nil {
		return nil, simplerr.Wrap(err, "can't insert")
	}

	return model.ToDomain(), nil
}

func (r *Repository) ListBadgesByUserIDs(ctx context.Context, userIDs []string) ([]*domain.Badge, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	query, args, err := sqlx.In("SELECT id, user_id, type, x, y, created_at FROM badge WHERE user_id IN(?)", userIDs)
	if err != nil {
		return nil, simplerr.Wrap(err, "can't build query")
	}

	var dest []*models.Badge

	if err := r.db.SelectContext(ctx, &dest, query, args...); err != nil {
		return nil, simplerr.Wrap(err, "can't select")
	}

	result := make([]*domain.Badge, len(dest))

	for idx, user := range dest {
		result[idx] = user.ToDomain()
	}

	return result, nil
}

func (r *Repository) UpdateBadge(ctx context.Context, badge *domain.Badge) (*domain.Badge, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	model := models.BadgeFromDomain(badge)

	query := `UPDATE badge SET x=:x, y=:y WHERE id = :id`

	if _, err := r.db.NamedExecContext(ctx, query, model); err != nil {
		return nil, simplerr.Wrap(err, "can't update")
	}

	return model.ToDomain(), nil
}
