package request

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	validation "github.com/gadavy/ozzo-validation/v4"

	"github.com/lissteron/cloudsuny/internal/app/codes"
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

	req.Bytes, err = ioutil.ReadAll(io.LimitReader(formData, maxImageSize))
	if err != nil {
		return nil, err
	}

	req.Format = http.DetectContentType(req.Bytes)

	if err := req.Validate(); err != nil {
		return nil, err
	}

	return req, nil
}

func (r *UploadImageRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Format, r.formatRules()...),
	)
}

func (r *UploadImageRequest) ToInput() *domain.Image {
	r.Format = strings.Split(r.Format, "/")[1]

	return &r.Image
}

func (r *UploadImageRequest) formatRules() []validation.Rule {
	return []validation.Rule{
		validation.In(
			"image/gif",
			"image/png",
			"image/jpeg",
		).ErrorCode(codes.ValidImageFormatIn.String()),
	}
}
