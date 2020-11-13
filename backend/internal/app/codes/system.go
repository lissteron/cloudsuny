package codes

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

// Validation error codes
type System int

func (s System) Int() int {
	return int(s)
}

func (s System) GRPC() int {
	return int(codes.Internal)
}

func (s System) HTTP() int {
	return http.StatusInternalServerError
}

const (
	DatabaseError System = 1000
	SystemError   System = 1001
)
