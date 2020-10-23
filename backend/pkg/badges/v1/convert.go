package badges

import (
	"github.com/lissteron/cloudsuny/internal/app/domain"
)

func (m *CreateReq) ToDomain() *domain.Badge {
	return &domain.Badge{
		UserID: m.UserId,
		Type:   m.Type,
		Point: domain.Point{
			X: m.GetPoint().GetX(),
			Y: m.GetPoint().GetY(),
		},
	}
}

func (m *UpdateReq) ToDomain() *domain.Badge {
	return &domain.Badge{
		ID: m.Id,
		Point: domain.Point{
			X: m.GetPoint().GetX(),
			Y: m.GetPoint().GetY(),
		},
	}
}