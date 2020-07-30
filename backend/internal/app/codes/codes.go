package codes

import (
	"net/http"
	"strconv"
)

const (
	system     = 1000
	useCase    = 2000
	validation = 200_000
)

const (
	DatabaseError = system + iota
	InvalidJSONError
)

// Validation error codes
type Code int

func (s Code) String() string {
	return strconv.Itoa(int(s))
}

const (
	ValidUserUsernameRequired Code = validation + iota
	ValidUserAvatarRequired
	ValidBadgeIDRequired
	ValidBadgeUserIDRequired
	ValidBadgeTypeRequired
	ValidBadgeTypeIn
)

func ToHTTP(code int) int {
	if code >= validation {
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}
