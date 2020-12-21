package codes

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

const (
	UserNotFound              Logic = 2000
	UserAlreadyExists         Logic = 2001
	BadgeNotFound             Logic = 2002
	InvalidUsernameOrPassword Logic = 2003
	SessionNotFound           Logic = 2004
	SessionExpired            Logic = 2005
)

type Logic int

func (s Logic) Int() int {
	return int(s)
}

func (s Logic) GRPC() int {
	switch s {
	case UserNotFound, BadgeNotFound:
		return int(codes.NotFound)
	case UserAlreadyExists:
		return int(codes.AlreadyExists)
	case InvalidUsernameOrPassword, SessionNotFound, SessionExpired:
		return int(codes.Unauthenticated)
	default:
		return int(codes.Unknown)
	}
}

func (s Logic) HTTP() int {
	switch s {
	case UserNotFound, BadgeNotFound:
		return http.StatusNotFound
	case UserAlreadyExists:
		return http.StatusConflict
	case InvalidUsernameOrPassword, SessionNotFound, SessionExpired:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
