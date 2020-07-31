package handlers

import (
	"context"
	"net/http"

	"github.com/lissteron/cloudsuny/internal/app/domain"
	"github.com/lissteron/cloudsuny/internal/app/http/request"
	"github.com/lissteron/cloudsuny/internal/app/http/response"
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

type CreateUser interface {
	Do(ctx context.Context, user *domain.User) (*domain.User, error)
}

type ViewUser interface {
	Do(ctx context.Context) ([]*domain.User, error)
}

type UserHandlers struct {
	create CreateUser
	view   ViewUser
	logger Logger
}

func NewUserHandlers(
	create CreateUser,
	view ViewUser,
	logger Logger,
) *UserHandlers {
	return &UserHandlers{
		create: create,
		view:   view,
		logger: logger,
	}
}

func (h *UserHandlers) CreateHandler(w http.ResponseWriter, r *http.Request) {
	resp, ctx := response.NewBaseResponse(), r.Context()
	defer resp.Write(w, h.logger)

	req, err := request.ReadCreateUserRequest(r)
	if err != nil {
		h.logger.Warnf("read create user request failed: %v", err)
		resp.ParseError(err)

		return
	}

	result, err := h.create.Do(ctx, req.ToInput())
	if err != nil {
		h.logger.Errorf("create user failed: %v", err)
		resp.ParseError(err)

		return
	}

	resp.SetData(result)
}

func (h *UserHandlers) ViewHandler(w http.ResponseWriter, r *http.Request) {
	resp, ctx := response.NewBaseResponse(), r.Context()
	defer resp.Write(w, h.logger)

	result, err := h.view.Do(ctx)
	if err != nil {
		h.logger.Errorf("view users failed: %v", err)
		resp.ParseError(err)

		return
	}

	resp.SetData(result)
}
