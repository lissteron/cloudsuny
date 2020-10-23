package http

import (
	"net/http"

	"github.com/loghole/tracing/tracelog"

	"github.com/lissteron/cloudsuny/internal/app/controllers/http/request"
	"github.com/lissteron/cloudsuny/internal/app/controllers/http/response"
	"github.com/lissteron/cloudsuny/internal/app/domain"
)

type ImagesService interface {
	UploadImage(img *domain.Image) (*domain.Image, error)
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
		h.logger.Warnf(ctx, "read upload image request failed: %v", err)
		resp.ParseError(err)

		return
	}

	result, err := h.service.UploadImage(req.ToInput())
	if err != nil {
		h.logger.Errorf(ctx, "upload image failed: %v", err)
		resp.ParseError(err)

		return
	}

	resp.SetData(result)
}

func (h *ImagesHandlers) FileServer() http.Handler {
	return http.FileServer(http.Dir(h.service.GetStoragePath()))
}
