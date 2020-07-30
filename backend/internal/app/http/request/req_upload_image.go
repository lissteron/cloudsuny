package request

import (
	"io"
	"io/ioutil"
	"net/http"
)

const (
	formImageData = "data"

	maxImageSize = 15 * 1024 * 1024 // 15 MB
)

type UploadImageRequest struct {
	Data []byte
}

func ReadUploadImageRequest(r *http.Request) (req *UploadImageRequest, err error) {
	req = &UploadImageRequest{}

	image, _, err := r.FormFile(formImageData)
	if err != nil {
		return nil, err
	}

	req.Data, err = ioutil.ReadAll(io.LimitReader(image, maxImageSize))
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (r *UploadImageRequest) ToInput() []byte {
	return r.Data
}
