package codes

import (
	"net/http"
	"strconv"

	"google.golang.org/grpc/codes"
)

const (
	ValidUserUsernameRequired Validation = 200_001
	ValidUserAvatarRequired   Validation = 200_002
	ValidBadgeIDRequired      Validation = 200_003
	ValidBadgeUserIDRequired  Validation = 200_004
	ValidBadgeTypeRequired    Validation = 200_005
	ValidBadgeTypeIn          Validation = 200_006
	ValidImageFormatIn        Validation = 200_007
	ValidUserIDRequired       Validation = 200_008
	ValidLimitMax             Validation = 200_009
	ValidLoginRequired        Validation = 200_010
	ValidPasswordRequired     Validation = 200_011
)

type Validation int

func (s Validation) Int() int {
	return int(s)
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
