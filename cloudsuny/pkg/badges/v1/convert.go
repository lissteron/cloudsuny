package badges

import (
	"github.com/lissteron/cloudsuny/internal/app/domain"
)

func (x *CreateReq) ToDomain() *domain.Badge {
	return &domain.Badge{
		UserID: x.UserId,
		Type:   x.Type,
		Point: domain.Point{
			X: x.GetPoint().GetX(),
			Y: x.GetPoint().GetY(),
		},
	}
}

func (x *UpdateReq) ToDomain() *domain.Badge {
	return &domain.Badge{
		ID: x.Id,
		Point: domain.Point{
			X: x.GetPoint().GetX(),
			Y: x.GetPoint().GetY(),
		},
	}
}
