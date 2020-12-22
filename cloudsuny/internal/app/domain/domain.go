package domain

import (
	"time"
)

const (
	SunBadge    = "sun"
	CloudBadge  = "cloud"
	IndianBadge = "indian"
)

type User struct {
	ID       string
	Username string
	Avatar   string
	Badges   []*Badge
}

type Badge struct {
	ID        string
	UserID    string
	Type      string
	Point     Point
	CreatedAt time.Time
}

type Point struct {
	X float64
	Y float64
}

type Image struct {
	Bytes  []byte `json:"-"`
	Format string `json:"format"`
	Name   string `json:"name"`
}

type UserListResult struct {
	List    []*User
	Total   int64
	Count   int64
	Skipped int64
}

func (r *UserListResult) UsersIDs() []string {
	result := make([]string, 0, len(r.List))

	for _, val := range r.List {
		result = append(result, val.ID)
	}

	return result
}
