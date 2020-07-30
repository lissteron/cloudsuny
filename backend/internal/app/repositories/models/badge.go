package models

import (
	"time"

	"github.com/lissteron/cloudsuny/internal/app/domain"
)

type Badge struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Type      string    `db:"type"`
	X         float64   `db:"x"`
	Y         float64   `db:"y"`
	CreatedAt time.Time `db:"created_at"`
}

func BadgeFromDomain(badge *domain.Badge) *Badge {
	return &Badge{
		ID:        badge.ID,
		UserID:    badge.UserID,
		Type:      badge.Type,
		X:         badge.Point.X,
		Y:         badge.Point.Y,
		CreatedAt: badge.CreatedAt,
	}
}

func (badge *Badge) ToDomain() *domain.Badge {
	return &domain.Badge{
		ID:        badge.ID,
		UserID:    badge.UserID,
		Type:      badge.Type,
		CreatedAt: badge.CreatedAt,
		Point: domain.Point{
			X: badge.X,
			Y: badge.Y,
		},
	}
}
