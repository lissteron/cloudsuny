package codes

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

const SystemError System = 1001

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
