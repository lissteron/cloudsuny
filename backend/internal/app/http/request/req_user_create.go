package request

import (
	"net/http"

	validation "github.com/gadavy/ozzo-validation/v4"
	json "github.com/json-iterator/go"
	"github.com/lissteron/simplerr"

	"github.com/lissteron/cloudsuny/internal/app/codes"
	"github.com/lissteron/cloudsuny/internal/app/domain"
)

type CreateUserRequest struct {
	domain.User
}

func ReadCreateUserRequest(r *http.Request) (*CreateUserRequest, error) {
	req := &CreateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return nil, simplerr.WrapWithCode(err, codes.InvalidJSONError, "invalid json struct")
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	return req, nil
}

func (r *CreateUserRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Username, validation.Required.ErrorCode(codes.ValidUserUsernameRequired.String())),
		validation.Field(&r.Avatar, validation.Required.ErrorCode(codes.ValidUserAvatarRequired.String())),
	)
}

func (r *CreateUserRequest) ToInput() *domain.User {
	return &r.User
}
