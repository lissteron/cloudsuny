package http

import (
	"context"
	"net/http"

	"github.com/loghole/tracing/tracelog"

	"github.com/lissteron/cloudsuny/internal/app/controllers/http/request"
	"github.com/lissteron/cloudsuny/internal/app/controllers/http/response"
	"github.com/lissteron/cloudsuny/internal/app/domain"
)

type ImagesService interface {
	UploadImage(ctx context.Context, img *domain.Image) (*domain.Image, error)
	GetStoragePath() string
}

type ImagesHandlers struct {
	service ImagesService
	logger  tracelog.Logger
}

func NewImagesHandlers(service ImagesService, logger tracelog.Logger) *ImagesHandlers {
	return &ImagesHandlers{
		service: service,
		logger:  logger,
	}
}

func (h *ImagesHandlers) UploadPhotoHandler(w http.ResponseWriter, r *http.Request) {
	ctx, resp := r.Context(), response.NewBaseResponse()
	defer resp.Write(w, h.logger)

	req, err := request.ReadUploadImageRequest(r)
	if err != nil {
		resp.ParseError(err)

		return
	}

	result, err := h.service.UploadImage(ctx, req.ToInput())
	if err != nil {
		resp.ParseError(err)

		return
	}

	resp.SetData(result)
}

func (h *ImagesHandlers) FileServer() http.Handler {
	return http.FileServer(http.Dir(h.service.GetStoragePath()))
}
