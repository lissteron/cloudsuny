package auth

import (
	validation "github.com/gadavy/ozzo-validation/v4"

	"github.com/lissteron/cloudsuny/internal/app/codes"
)

func (x *SignInReq) Validate() error {
	return validation.ValidateStruct(x,
		validation.Field(&x.Login, validation.Required.ErrorCode(codes.ValidLoginRequired.String())),
		validation.Field(&x.Password, validation.Required.ErrorCode(codes.ValidPasswordRequired.String())),
	)
}
