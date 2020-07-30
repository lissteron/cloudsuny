package interfaces

import (
	"context"

	"github.com/lissteron/cloudsuny/internal/app/domain"
	"github.com/lissteron/cloudsuny/internal/app/repositories"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
}

type Storage interface {
	WithTransaction(ctx context.Context) (*repositories.Repository, error)
	CommitTransaction() error
	RollbackTransaction()
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	ListUsers(ctx context.Context) ([]*domain.User, error)
	FindUserByID(ctx context.Context, userID string) (*domain.User, error)
	FindUserByUsername(ctx context.Context, username string) (*domain.User, error)

	CreateBadge(ctx context.Context, badge *domain.Badge) (*domain.Badge, error)
	FindBadgeByID(ctx context.Context, badgeID string) (*domain.Badge, error)
	ListBadges(ctx context.Context) ([]*domain.Badge, error)
	UpdateBadge(ctx context.Context, badge *domain.Badge) (*domain.Badge, error)
}
