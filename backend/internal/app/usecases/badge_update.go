package usecases

import (
	"context"
	"errors"

	"github.com/lissteron/cloudsuny/internal/app/domain"
	"github.com/lissteron/cloudsuny/internal/app/usecases/interfaces"
)

var (
	ErrBadgeNotFound = errors.New("badge not found")
)

type UpdateBadge struct {
	storage interfaces.Storage
	logger  interfaces.Logger
}

func NewUpdateBadge(
	storage interfaces.Storage,
	logger interfaces.Logger,
) *UpdateBadge {
	return &UpdateBadge{
		storage: storage,
		logger:  logger,
	}
}

func (c *UpdateBadge) Do(ctx context.Context, badge *domain.Badge) (*domain.Badge, error) {
	storage, err := c.storage.WithTransaction(ctx)
	if err != nil {
		c.logger.Errorf("start tx failed: %v", err)
		return nil, err
	}

	defer storage.RollbackTransaction()

	exists, err := storage.FindBadgeByID(ctx, badge.ID)
	if err != nil {
		c.logger.Errorf("find badge failed: %v", err)
		return nil, err
	}

	if exists == nil {
		c.logger.Errorf("badge with id = %s not found", badge.ID)
		return nil, ErrBadgeNotFound
	}

	exists.Point = badge.Point

	badge, err = storage.UpdateBadge(ctx, exists)
	if err != nil {
		c.logger.Errorf("update badge failed: %v", err)
		return nil, err
	}

	if err := storage.CommitTransaction(); err != nil {
		c.logger.Errorf("commit tx failed: %v", err)
		return nil, err
	}

	return badge, nil
}
