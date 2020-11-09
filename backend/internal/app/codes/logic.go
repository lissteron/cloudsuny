package codes

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

const logic = 2000

type Logic int

func (s Logic) Int() int {
	return int(s) + logic
}

func (s Logic) GRPC() int {
	switch s {
	case UserNotFound, BadgeNotFound:
		return int(codes.NotFound)
	case UserAlreadyExists:
		return int(codes.AlreadyExists)
	}

	return int(codes.Unknown)
}

func (s Logic) HTTP() int {
	switch s {
	case UserNotFound, BadgeNotFound:
		return http.StatusNotFound
	case UserAlreadyExists:
		return http.StatusConflict
	}

	return http.StatusInternalServerError
}

const (
	UserNotFound Logic = iota + 1
	UserAlreadyExists
	BadgeNotFound
)
