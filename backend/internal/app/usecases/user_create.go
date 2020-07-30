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
	ErrUserAlreadyExists = errors.New("user already exists")
)

type CreateUser struct {
	storage interfaces.Storage
	logger  interfaces.Logger
}

func NewCreateUser(
	storage interfaces.Storage,
	logger interfaces.Logger,
) *CreateUser {
	return &CreateUser{
		storage: storage,
		logger:  logger,
	}
}

func (c *CreateUser) Do(ctx context.Context, user *domain.User) (*domain.User, error) {
	storage, err := c.storage.WithTransaction(ctx)
	if err != nil {
		c.logger.Errorf("start tx failed: %v", err)
		return nil, simplerr.WithCode(err, codes.DatabaseError)
	}

	defer storage.RollbackTransaction()

	exists, err := storage.FindUserByUsername(ctx, user.Username)
	if err != nil {
		c.logger.Errorf("find user failed: %v", err)
		return nil, simplerr.WithCode(err, codes.DatabaseError)
	}

	if exists != nil {
		return nil, simplerr.WithCode(ErrUserAlreadyExists, codes.UserAlreadyExists)
	}

	user, err = storage.CreateUser(ctx, user)
	if err != nil {
		c.logger.Errorf("create user failed: %v", err)
		return nil, simplerr.WithCode(err, codes.DatabaseError)
	}

	if err := storage.CommitTransaction(); err != nil {
		c.logger.Errorf("commit tx failed: %v", err)
		return nil, simplerr.WithCode(err, codes.DatabaseError)
	}

	return user, nil
}
