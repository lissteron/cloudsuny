package request

import (
	"net/http"

	validation "github.com/gadavy/ozzo-validation/v4"
	jsoniter "github.com/json-iterator/go"
	"github.com/lissteron/simplerr"

	"github.com/lissteron/cloudsuny/internal/app/codes"
	"github.com/lissteron/cloudsuny/internal/app/domain"
)

type CreateBadgeRequest struct {
	domain.Badge
}

func ReadCreateBadgeRequest(r *http.Request) (*CreateBadgeRequest, error) {
	req := &CreateBadgeRequest{}

	if err := jsoniter.NewDecoder(r.Body).Decode(req); err != nil {
		return nil, simplerr.WrapWithCode(err, codes.InvalidJSONError, "invalid json struct")
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	return req, nil
}
func (r *CreateBadgeRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.UserID, r.userIDRules()...),
		validation.Field(&r.Type, r.typeRules()...),
	)
}

func (r *CreateBadgeRequest) ToInput() *domain.Badge {
	return &r.Badge
}

func (r *CreateBadgeRequest) userIDRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidBadgeUserIDRequired.String()),
	}
}

func (r *CreateBadgeRequest) typeRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidBadgeTypeRequired.String()),
		validation.In(domain.CloudBadge, domain.SunBadge).ErrorCode(codes.ValidBadgeTypeIn.String()),
	}
}
