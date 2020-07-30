package usecases

import (
	"context"

	"github.com/lissteron/cloudsuny/internal/app/domain"
	"github.com/lissteron/cloudsuny/internal/app/usecases/interfaces"
)

type ListUser struct {
	storage interfaces.Storage
	logger  interfaces.Logger
}

func NewListUser(
	storage interfaces.Storage,
	logger interfaces.Logger,
) *ListUser {
	return &ListUser{
		storage: storage,
		logger:  logger,
	}
}

func (l *ListUser) Do(ctx context.Context) ([]*domain.User, error) {
	users, err := l.storage.ListUsers(ctx)
	if err != nil {
		l.logger.Errorf("list users failed: %v", err)
		return nil, err
	}

	return users, nil
}
