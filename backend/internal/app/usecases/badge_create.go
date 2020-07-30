package usecases

import (
	"context"
	"errors"

	"github.com/lissteron/cloudsuny/internal/app/domain"
	"github.com/lissteron/cloudsuny/internal/app/usecases/interfaces"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type CreateBadge struct {
	storage interfaces.Storage
	logger  interfaces.Logger
}

func NewCreateBadge(
	storage interfaces.Storage,
	logger interfaces.Logger,
) *CreateBadge {
	return &CreateBadge{
		storage: storage,
		logger:  logger,
	}
}

func (c *CreateBadge) Do(ctx context.Context, badge *domain.Badge) (*domain.Badge, error) {
	user, err := c.storage.FindUserByID(ctx, badge.UserID)
	if err != nil {
		c.logger.Errorf("find user by id failed: %v", err)
		return nil, err
	}

	if user == nil {
		c.logger.Error("user not found")
		return nil, ErrUserNotFound
	}

	badge, err = c.storage.CreateBadge(ctx, badge)
	if err != nil {
		c.logger.Errorf("create badge failed: %v", err)
		return nil, err
	}

	return badge, nil
}