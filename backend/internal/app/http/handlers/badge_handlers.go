package handlers

import (
	"context"
	"net/http"

	"github.com/lissteron/cloudsuny/internal/app/domain"
	"github.com/lissteron/cloudsuny/internal/app/http/request"
	"github.com/lissteron/cloudsuny/internal/app/http/response"
)

type CreateBadge interface {
	Do(ctx context.Context, badge *domain.Badge) (*domain.Badge, error)
}

type UpdateBadge interface {
	Do(ctx context.Context, badge *domain.Badge) (*domain.Badge, error)
}

type BadgeHandlers struct {
	create CreateBadge
	update UpdateBadge
	logger Logger
}

func NewBadgeHandlers(
	create CreateBadge,
	update UpdateBadge,
	logger Logger,
) *BadgeHandlers {
	return &BadgeHandlers{
		create: create,
		update: update,
		logger: logger,
	}
}

func (h *BadgeHandlers) CreateHandler(w http.ResponseWriter, r *http.Request) {
	resp, ctx := response.NewBaseResponse(), r.Context()
	defer resp.Write(w, h.logger)

	req, err := request.ReadCreateBadgeRequest(r)
	if err != nil {
		h.logger.Errorf("read create badge request failed: %v", err)
		resp.ParseError(err)

		return
	}

	result, err := h.create.Do(ctx, req.ToInput())
	if err != nil {
		h.logger.Errorf("create badge failed: %v", err)
		resp.ParseError(err)

		return
	}

	resp.SetData(result)
}

func (h *BadgeHandlers) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	resp, ctx := response.NewBaseResponse(), r.Context()
	defer resp.Write(w, h.logger)

	req, err := request.ReadUpdateBadgeRequest(r)
	if err != nil {
		h.logger.Errorf("read update badge request failed: %v", err)
		resp.ParseError(err)

		return
	}

	result, err := h.update.Do(ctx, req.ToInput())
	if err != nil {
		h.logger.Errorf("update badge failed: %v", err)
		resp.ParseError(err)

		return
	}

	resp.SetData(result)
}
