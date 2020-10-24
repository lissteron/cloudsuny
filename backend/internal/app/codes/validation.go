package codes

import (
	"net/http"
	"strconv"

	"google.golang.org/grpc/codes"
)

const validation = 200_000

type Validation int

func (s Validation) Int() int {
	return int(s) + validation
}

func (s Validation) GRPC() int {
	return int(codes.InvalidArgument)
}

func (s Validation) HTTP() int {
	return http.StatusBadRequest
}

func (s Validation) String() string {
	return strconv.Itoa(int(s))
}

const (
	ValidationErr Validation = iota + 1
	ValidUserUsernameRequired
	ValidUserAvatarRequired
	ValidBadgeIDRequired
	ValidBadgeUserIDRequired
	ValidBadgeTypeRequired
	ValidBadgeTypeIn
	ValidImageFormatIn
)
