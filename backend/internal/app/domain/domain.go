package domain

import (
	"time"
)

type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Avatar   string   `json:"avatar"`
	Badges   []*Badge `json:"badges"`
}

const (
	SunBadge   = "sun"
	CloudBadge = "cloud"
)

type Badge struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Type      string    `json:"type"`
	Point     Point     `json:"point"`
	CreatedAt time.Time `json:"created_at"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
