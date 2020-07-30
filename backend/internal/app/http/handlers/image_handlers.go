package handlers

import (
	"net/http"

	"github.com/lissteron/cloudsuny/internal/app/domain"
	"github.com/lissteron/cloudsuny/internal/app/http/request"
	"github.com/lissteron/cloudsuny/internal/app/http/response"
)

type ImagesService interface {
	UploadImage(img *domain.Image) (*domain.Image, error)
	GetStoragePath() string
}

type ImagesHandlers struct {
	service ImagesService
	logger  Logger
}

func NewImagesHandlers(service ImagesService, logger Logger) *ImagesHandlers {
	return &ImagesHandlers{
		service: service,
		logger:  logger,
	}
}

func (h *ImagesHandlers) UploadPhotoHandler(w http.ResponseWriter, r *http.Request) {
	resp := response.NewBaseResponse()
	defer resp.Write(w, h.logger)

	req, err := request.ReadUploadImageRequest(r)
	if err != nil {
		h.logger.Errorf("read upload image request failed: %v", err)
		resp.ParseError(err)

		return
	}

	result, err := h.service.UploadImage(req.ToInput())
	if err != nil {
		h.logger.Errorf("upload image failed: %v", err)
		resp.ParseError(err)

		return
	}

	resp.SetData(result)
}

func (h *ImagesHandlers) FileServer() http.Handler {
	return http.FileServer(http.Dir(h.service.GetStoragePath()))
}
