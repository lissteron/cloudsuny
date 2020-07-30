package usecases

import (
	"context"
	"errors"

	"github.com/lissteron/simplerr"

	"github.com/lissteron/cloudsuny/internal/app/codes"
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
		return nil, simplerr.WrapWithCode(err, codes.DatabaseError, "find user by id failed")
	}

	if user == nil {
		c.logger.Error("user not found")
		return nil, simplerr.WrapWithCode(ErrUserNotFound, codes.UserNotFound, "user not found")
	}

	badge, err = c.storage.CreateBadge(ctx, badge)
	if err != nil {
		c.logger.Errorf("create badge failed: %v", err)
		return nil, simplerr.WrapWithCode(err, codes.DatabaseError, "create badge failed")
	}

	return badge, nil
}
