package users

import (
	validation "github.com/gadavy/ozzo-validation/v4"

	"github.com/lissteron/cloudsuny/internal/app/codes"
)

func (m *CreateReq) Validate() error {
	return validation.ValidateStruct(m,
		validation.Field(&m.Username, validation.Required.ErrorCode(codes.ValidUserUsernameRequired.String())),
		validation.Field(&m.Avatar, validation.Required.ErrorCode(codes.ValidUserAvatarRequired.String())),
	)
}
