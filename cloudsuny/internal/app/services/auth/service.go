package auth

import (
	"context"
	"encoding/hex"
	"errors"
	"io"
	"math/rand"
	"time"

	"github.com/lissteron/simplerr"
	"github.com/loghole/tracing"
	"github.com/loghole/tracing/tracelog"

	"github.com/lissteron/cloudsuny/config"
	"github.com/lissteron/cloudsuny/internal/app/codes"
	"github.com/lissteron/cloudsuny/internal/app/domain"
)

const keySize = 32

var (
	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
	ErrExpiredToken              = errors.New("expired session")
)

type Storage interface {
	AddSession(ctx context.Context, session *domain.Session) (*domain.Session, error)
	FindSession(ctx context.Context, token string) (*domain.Session, error)
	UpdateSession(ctx context.Context, session *domain.Session) error
	DeleteSession(ctx context.Context, token string) error
}

type Service struct {
	storage Storage
	logger  tracelog.Logger
	rnd     io.Reader
}

func NewService(storage Storage, logger tracelog.Logger) *Service {
	return &Service{
		storage: storage,
		logger:  logger,
		rnd:     rand.New(rand.NewSource(time.Now().Unix())), // nolint:gosec // нет смысла грузить нормальный рандом
	}
}

func (s *Service) SignIn(ctx context.Context, login, password string) (*domain.Session, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	existsPassword, ok := config.UserCredentials()[login]
	if !ok || existsPassword != password {
		return nil, simplerr.WithCode(ErrInvalidUsernameOrPassword, codes.InvalidUsernameOrPassword)
	}

	token, err := s.buildToken()
	if err != nil {
		s.logger.Error(ctx, err)

		return nil, simplerr.Wrap(err, "can't build token")
	}

	session, err := s.storage.AddSession(ctx, &domain.Session{Login: login, Token: token})
	if err != nil {
		s.logger.Error(ctx, err)

		return nil, simplerr.Wrap(err, "can't store token")
	}

	return session, nil
}

func (s *Service) SignOut(ctx context.Context, token string) error {
	defer tracing.ChildSpan(&ctx).Finish()

	session, err := s.Authenticate(ctx, token)
	if err != nil {
		return simplerr.Wrap(err, "can't authenticate")
	}

	if err := s.storage.DeleteSession(ctx, session.Token); err != nil {
		s.logger.Error(ctx, err)

		return simplerr.Wrap(err, "can't drop session")
	}

	return nil
}

func (s *Service) Authenticate(ctx context.Context, token string) (*domain.Session, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	session, err := s.storage.FindSession(ctx, token)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, simplerr.WithCode(err, codes.SessionNotFound)
		}

		s.logger.Error(ctx, err)

		return nil, simplerr.Wrap(err, "can't find session")
	}

	if session.IsExpired(config.GetSessionMaxAge()) {
		if err := s.storage.DeleteSession(ctx, session.Token); err != nil {
			s.logger.Error(ctx, err)

			return nil, simplerr.Wrap(err, "can't drop session")
		}

		return nil, simplerr.WithCode(ErrExpiredToken, codes.SessionExpired)
	}

	if err := s.storage.UpdateSession(ctx, session); err != nil {
		s.logger.Error(ctx, err)

		return nil, simplerr.Wrap(err, "can't update session")
	}

	return session, nil
}

// Так себе токен, для данного приложения сойдет, но если
// двигаться дальше - необходимо заменить на нормальную генерацию.
func (s *Service) buildToken() (string, error) {
	buf := make([]byte, keySize)

	if _, err := s.rnd.Read(buf); err != nil {
		return "", simplerr.Wrap(err, "can't read rand bytes")
	}

	return hex.EncodeToString(buf), nil
}
