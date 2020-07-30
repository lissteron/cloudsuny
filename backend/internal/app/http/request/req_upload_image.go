package request

import (
	"image"
	_ "image/gif"  // gif package
	_ "image/jpeg" // jpeg package
	_ "image/png"  // png package
	"io"
	"net/http"

	"github.com/lissteron/cloudsuny/internal/app/domain"
)

const (
	formImageData = "data"

	maxImageSize = 15 * 1024 * 1024 // 15 MB
)

type UploadImageRequest struct {
	domain.Image
}

func ReadUploadImageRequest(r *http.Request) (req *UploadImageRequest, err error) {
	req = &UploadImageRequest{}

	formData, _, err := r.FormFile(formImageData)
	if err != nil {
		return nil, err
	}

	req.Image.Image, req.Format, err = image.Decode(io.LimitReader(formData, maxImageSize))
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (r *UploadImageRequest) ToInput() *domain.Image {
	return &r.Image
}
