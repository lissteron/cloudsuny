package codes

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

const system = 1000

// Validation error codes
type System int

func (s System) Int() int {
	return int(s) + system
}

func (s System) GRPC() int {
	return int(codes.Internal)
}

func (s System) HTTP() int {
	return http.StatusInternalServerError
}

const (
	DatabaseError System = iota + 1
	SystemError
)
