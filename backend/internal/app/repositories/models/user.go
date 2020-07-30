package models

import (
	"time"

	"github.com/lissteron/cloudsuny/internal/app/domain"
)

type User struct {
	ID        string    `db:"id"`
	Username  string    `db:"username"`
	Avatar    string    `db:"avatar"`
	CreatedAt time.Time `db:"created_at"`
}

func UserFromDomain(user *domain.User) *User {
	return &User{
		ID:       user.ID,
		Username: user.Username,
		Avatar:   user.Avatar,
	}
}

func (user *User) ToDomain() *domain.User {
	return &domain.User{
		ID:       user.ID,
		Username: user.Username,
		Avatar:   user.Avatar,
	}
}
