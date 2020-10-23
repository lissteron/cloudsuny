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
	return int(codes.InvalidArgument)
}

func (s Logic) HTTP() int {
	return http.StatusBadRequest
}

const (
	UserNotFound Logic = iota + 1
	UserAlreadyExists
	BadgeNotFound
	UnknownImgFormat
)
