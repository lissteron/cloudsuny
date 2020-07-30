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
	SystemError
	InvalidJSONError
	EncodeImgFailed
)

const (
	UserNotFound = useCase + iota
	UserAlreadyExists
	BadgeNotFound
	UnknownImgFormat
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

	switch code {
	case UserNotFound, BadgeNotFound:
		return http.StatusNotFound
	case UserAlreadyExists:
		return http.StatusConflict
	case UnknownImgFormat:
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}
