package badges

import (
	"context"
	"errors"

	"github.com/lissteron/simplerr"
	"github.com/loghole/tracing"
	"github.com/loghole/tracing/tracelog"

	"github.com/lissteron/cloudsuny/internal/app/codes"
	"github.com/lissteron/cloudsuny/internal/app/domain"
)

type Storage interface {
	FindUserByID(ctx context.Context, userID string) (*domain.User, error)
	CreateBadge(ctx context.Context, badge *domain.Badge) (*domain.Badge, error)
	UpdateBadge(ctx context.Context, badge *domain.Badge) (*domain.Badge, error)
}

type Service struct {
	storage Storage
	logger  tracelog.Logger
}

func NewService(storage Storage, logger tracelog.Logger) *Service {
	return &Service{
		storage: storage,
		logger:  logger,
	}
}

func (s *Service) Create(ctx context.Context, badge *domain.Badge) (*domain.Badge, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	if _, err := s.storage.FindUserByID(ctx, badge.UserID); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, simplerr.WrapfWithCode(err, codes.UserNotFound, "user '%s' not found", badge.UserID)
		}

		s.logger.Errorf(ctx, "find user by id failed: %v", err)

		return nil, simplerr.Wrap(err, "find user failed")
	}

	badge, err := s.storage.CreateBadge(ctx, badge)
	if err != nil {
		s.logger.Errorf(ctx, "create badge failed: %v", err)

		return nil, simplerr.Wrap(err, "create badge failed")
	}

	return badge, nil
}

func (s *Service) Update(ctx context.Context, badge *domain.Badge) (*domain.Badge, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	badge, err := s.storage.UpdateBadge(ctx, badge)
	if err != nil {
		s.logger.Errorf(ctx, "update badge failed: %v", err)

		return nil, simplerr.Wrap(err, "update badge failed")
	}

	return badge, nil
}
