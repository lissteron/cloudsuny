package users

import (
	"github.com/lissteron/cloudsuny/internal/app/domain"
	types "github.com/lissteron/cloudsuny/pkg/types/v1"
)

func UserFromDomain(user *domain.User) *User {
	output := &User{
		Id:       user.ID,
		Username: user.Username,
		Avatar:   user.Avatar,
	}

	for _, badge := range user.Badges {
		output.Badges = append(output.Badges, types.BadgeFromDomain(badge))
	}

	return output
}

func (m *CreateUsersReq) ToDomain() *domain.User {
	return &domain.User{
		Username: m.Username,
		Avatar:   m.Avatar,
	}
}
