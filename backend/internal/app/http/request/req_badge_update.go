package request

import (
	"net/http"

	validation "github.com/gadavy/ozzo-validation/v4"
	jsoniter "github.com/json-iterator/go"
	"github.com/lissteron/simplerr"

	"github.com/lissteron/cloudsuny/internal/app/codes"
	"github.com/lissteron/cloudsuny/internal/app/domain"
)

type UpdateBadgeRequest struct {
	domain.Badge
}

func ReadUpdateBadgeRequest(r *http.Request) (*UpdateBadgeRequest, error) {
	req := &UpdateBadgeRequest{}

	if err := jsoniter.NewDecoder(r.Body).Decode(req); err != nil {
		return nil, simplerr.WrapWithCode(err, codes.InvalidJSONError, "invalid json struct")
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	return req, nil
}
func (r *UpdateBadgeRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.ID, r.idRules()...),
	)
}

func (r *UpdateBadgeRequest) ToInput() *domain.Badge {
	return &r.Badge
}

func (r *UpdateBadgeRequest) idRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidBadgeIDRequired.String()),
	}
}
