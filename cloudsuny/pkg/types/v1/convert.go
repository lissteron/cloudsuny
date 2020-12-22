package types

import (
	"github.com/gadavy/pb-types/timestamp"

	"github.com/lissteron/cloudsuny/internal/app/domain"
)

func BadgeFromDomain(input *domain.Badge) *Badge {
	badge := &Badge{
		Id:     input.ID,
		UserId: input.UserID,
		Type:   input.Type,
		Point: &Point{
			X: input.Point.X,
			Y: input.Point.Y,
		},
		CreatedAt: timestamp.TimeToProto(input.CreatedAt),
	}

	return badge
}
