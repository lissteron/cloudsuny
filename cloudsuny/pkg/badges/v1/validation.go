package badges

import (
	validation "github.com/gadavy/ozzo-validation/v4"

	"github.com/lissteron/cloudsuny/internal/app/codes"
	"github.com/lissteron/cloudsuny/internal/app/domain"
)

func (x *CreateReq) Validate() error {
	return validation.ValidateStruct(x,
		validation.Field(&x.UserId, x.userIDRules()...),
		validation.Field(&x.Type, x.typeRules()...),
	)
}

func (x *CreateReq) userIDRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidBadgeUserIDRequired.String()),
	}
}

func (x *CreateReq) typeRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidBadgeTypeRequired.String()),
		validation.In(domain.CloudBadge, domain.SunBadge, domain.IndianBadge).ErrorCode(codes.ValidBadgeTypeIn.String()),
	}
}

func (x *UpdateReq) Validate() error {
	return validation.ValidateStruct(x,
		validation.Field(&x.Id, x.idRules()...),
	)
}

func (x *UpdateReq) idRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidBadgeIDRequired.String()),
	}
}
