package users

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
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	ListUsers(ctx context.Context, limit, offset uint64) (*domain.UserListResult, error)
	ListBadgesByUserIDs(ctx context.Context, userIDs []string) ([]*domain.Badge, error)
	DeleteUser(ctx context.Context, userID string) error
}

type Service struct {
	logger  tracelog.Logger
	storage Storage
}

func NewService(storage Storage, logger tracelog.Logger) *Service {
	return &Service{
		storage: storage,
		logger:  logger,
	}
}

func (s *Service) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	user, err := s.storage.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, domain.ErrAlreadyExists) {
			return nil, simplerr.WrapWithCode(err, codes.UserAlreadyExists, "user already exists")
		}

		s.logger.Errorf(ctx, "create user failed: %v", err)

		return nil, simplerr.Wrap(err, "create user failed")
	}

	return user, nil
}

func (s *Service) List(ctx context.Context, limit, offset uint64) (*domain.UserListResult, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	users, err := s.storage.ListUsers(ctx, limit, offset)
	if err != nil {
		s.logger.Errorf(ctx, "list users failed: %v", err)

		return nil, simplerr.Wrap(err, "list users failed")
	}

	if users.Count == 0 {
		return users, nil
	}

	var (
		index = make(map[string]*domain.User, users.Count)
		ids   = make([]string, 0, users.Count)
	)

	for _, user := range users.List {
		ids, index[user.ID] = append(ids, user.ID), user
	}

	badges, err := s.storage.ListBadgesByUserIDs(ctx, ids)
	if err != nil {
		s.logger.Errorf(ctx, "list badges failed: %v", err)

		return nil, simplerr.Wrap(err, "list badges failed")
	}

	for _, badge := range badges {
		if user, ok := index[badge.UserID]; ok {
			user.Badges = append(user.Badges, badge)
		}
	}

	return users, nil
}

func (s *Service) Delete(ctx context.Context, userID string) error {
	defer tracing.ChildSpan(&ctx).Finish()

	if err := s.storage.DeleteUser(ctx, userID); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return simplerr.WrapfWithCode(err, codes.UserNotFound, "user '%s' not found", userID)
		}

		s.logger.Errorf(ctx, "delete user failed: %v", err)

		return simplerr.Wrap(err, "delete user failed")
	}

	return nil
}
