package users

import (
	"github.com/lissteron/cloudsuny/internal/app/domain"
	types "github.com/lissteron/cloudsuny/pkg/types/v1"
)

func UserListFromDomain(list []*domain.User) []*User {
	result := make([]*User, 0, len(list))

	for _, val := range list {
		result = append(result, UserFromDomain(val))
	}

	return result
}

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

func (x *CreateReq) ToDomain() *domain.User {
	return &domain.User{
		Username: x.Username,
		Avatar:   x.Avatar,
	}
}
