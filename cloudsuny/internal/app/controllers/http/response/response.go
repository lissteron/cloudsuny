package response

import (
	"context"
	"net/http"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/lissteron/simplerr"
)

type Logger interface {
	Errorf(ctx context.Context, template string, args ...interface{})
}

type BaseResponse struct {
	Status int         `json:"-"`
	Errors []RespError `json:"errors"`
	Data   interface{} `json:"data"`
}

type RespError struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func NewBaseResponse() *BaseResponse {
	return &BaseResponse{Status: http.StatusOK}
}

func (r *BaseResponse) Write(w http.ResponseWriter, log Logger) {
	w.Header().Add("Content-Type", "application/json")

	if r.Status != 0 {
		w.WriteHeader(r.Status)
	}

	if err := jsoniter.NewEncoder(w).Encode(r); err != nil {
		log.Errorf(context.TODO(), "write response failed: %v", err)
	}
}

func (r *BaseResponse) SetData(v interface{}) {
	r.Data = v
}

func (r *BaseResponse) ParseError(err error) {
	r.Status = simplerr.GetCode(err).HTTP()

	r.Errors = append(r.Errors, RespError{
		Code:   strconv.Itoa(simplerr.GetCode(err).Int()),
		Detail: err.Error(),
	})
}
