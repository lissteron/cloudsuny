package domain

import (
	"time"
)

type Session struct {
	Login     string    `db:"login"`
	Token     string    `db:"token"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (s *Session) IsExpired(maxAge time.Duration) bool {
	return time.Now().After(s.UpdatedAt.Add(maxAge))
}
