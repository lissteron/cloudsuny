package domain

import (
	"time"
)

type User struct {
	ID       string   `json:"id,omitempty"`
	Username string   `json:"username,omitempty"`
	Avatar   string   `json:"avatar,omitempty"`
	Badges   []*Badge `json:"badges,omitempty"`
}

const (
	SunBadge   = "sun"
	CloudBadge = "cloud"
)

type Badge struct {
	ID        string    `json:"id,omitempty"`
	UserID    string    `json:"user_id,omitempty"`
	Type      string    `json:"type,omitempty"`
	Point     Point     `json:"point,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Image struct {
	Bytes  []byte `json:"-"`
	Format string `json:"-"`
	Name   string `json:"name"`
}
