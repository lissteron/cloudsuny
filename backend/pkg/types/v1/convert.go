package types

import (
	"github.com/lissteron/cloudsuny/internal/app/domain"
	"github.com/lissteron/cloudsuny/pkg/prototime"
)

func BadgeFromDomain(input *domain.Badge) *Badge {
	return &Badge{
		Id:     input.ID,
		UserId: input.UserID,
		Type:   input.Type,
		Point: &Point{
			X: input.Point.X,
			Y: input.Point.Y,
		},
		CreatedAt: prototime.TimeToProto(input.CreatedAt),
	}
}
