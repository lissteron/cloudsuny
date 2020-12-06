package types

import (
	"github.com/gogo/protobuf/types"

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
	}

	badge.CreatedAt, _ = types.TimestampProto(input.CreatedAt)

	return badge
}
