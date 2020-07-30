package usecases

import (
	"context"

	"github.com/lissteron/simplerr"

	"github.com/lissteron/cloudsuny/internal/app/codes"
	"github.com/lissteron/cloudsuny/internal/app/domain"
	"github.com/lissteron/cloudsuny/internal/app/usecases/interfaces"
)

type ViewUser struct {
	storage interfaces.Storage
	logger  interfaces.Logger
}

func NewViewUser(
	storage interfaces.Storage,
	logger interfaces.Logger,
) *ViewUser {
	return &ViewUser{
		storage: storage,
		logger:  logger,
	}
}

func (l *ViewUser) Do(ctx context.Context) ([]*domain.User, error) {
	users, err := l.storage.ListUsers(ctx)
	if err != nil {
		l.logger.Errorf("list users failed: %v", err)
		return nil, simplerr.WrapWithCode(err, codes.DatabaseError, "list users failed")
	}

	index := make(map[string]*domain.User, len(users))

	for _, user := range users {
		index[user.ID] = user
	}

	badges, err := l.storage.ListBadges(ctx)
	if err != nil {
		l.logger.Errorf("list badges failed: %v", err)
		return nil, simplerr.WrapWithCode(err, codes.DatabaseError, "list badges failed")
	}

	for _, badge := range badges {
		if user, ok := index[badge.UserID]; ok {
			user.Badges = append(user.Badges, badge)
		}
	}

	return users, nil
}
