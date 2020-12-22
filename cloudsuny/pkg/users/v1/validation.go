package users

import (
	validation "github.com/gadavy/ozzo-validation/v4"

	"github.com/lissteron/cloudsuny/internal/app/codes"
)

const maxLimit uint64 = 100

func (x *CreateReq) Validate() error {
	return validation.ValidateStruct(x,
		validation.Field(&x.Username, validation.Required.ErrorCode(codes.ValidUserUsernameRequired.String())),
		validation.Field(&x.Avatar, validation.Required.ErrorCode(codes.ValidUserAvatarRequired.String())),
	)
}

func (x *DeleteReq) Validate() error {
	return validation.ValidateStruct(x,
		validation.Field(&x.UserId, validation.Required.ErrorCode(codes.ValidUserIDRequired.String())),
	)
}

func (x *ListWithBadgesReq) Validate() error {
	return validation.ValidateStruct(x,
		validation.Field(&x.Limit, validation.Max(maxLimit).ErrorCode(codes.ValidLimitMax.String())),
	)
}
