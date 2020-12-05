package badges

import (
	validation "github.com/gadavy/ozzo-validation/v4"
	"github.com/lissteron/simplerr"

	"github.com/lissteron/cloudsuny/internal/app/codes"
	"github.com/lissteron/cloudsuny/internal/app/domain"
)

func (m *CreateReq) Validate() error {
	err := validation.ValidateStruct(m,
		validation.Field(&m.UserId, m.userIDRules()...),
		validation.Field(&m.Type, m.typeRules()...),
	)
	if err != nil {
		return simplerr.WithCode(err, codes.ValidationErr)
	}

	return nil
}

func (m *CreateReq) userIDRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidBadgeUserIDRequired.String()),
	}
}

func (m *CreateReq) typeRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidBadgeTypeRequired.String()),
		validation.In(domain.CloudBadge, domain.SunBadge, domain.IndianBadge).ErrorCode(codes.ValidBadgeTypeIn.String()),
	}
}

func (m *UpdateReq) Validate() error {
	return validation.ValidateStruct(m,
		validation.Field(&m.Id, m.idRules()...),
	)
}

func (m *UpdateReq) idRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidBadgeIDRequired.String()),
	}
}
