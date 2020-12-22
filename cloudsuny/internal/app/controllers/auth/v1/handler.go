// Generated by tron. You must modify it.
package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/lissteron/simplerr"
	"github.com/loghole/tracing/tracelog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/lissteron/cloudsuny/internal/app/domain"
	authV1 "github.com/lissteron/cloudsuny/pkg/auth/v1"

	transport "github.com/loghole/tron/transport"
)

const (
	cookieName   = "_sess_1"
	cookieHeader = "cookie"
	headerParts  = 2
	headerSep    = "="
)

type Service interface {
	SignIn(ctx context.Context, login, password string) (*domain.Session, error)
	SignOut(ctx context.Context, token string) error
	Authenticate(ctx context.Context, token string) (*domain.Session, error)
}

type Implementation struct {
	authV1.UnimplementedAuthenticationServer
	service Service
	logger  tracelog.Logger
}

func NewImplementation(service Service, logger tracelog.Logger) *Implementation {
	return &Implementation{
		service: service,
		logger:  logger,
	}
}

// GetDescription is a simple alias to the ServiceDesc constructor.
// It makes it possible to register the service implementation @ the server.
func (i *Implementation) GetDescription() transport.ServiceDesc {
	return authV1.NewAuthenticationServiceDesc(i)
}

func (i *Implementation) readGRPCCookie(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	for _, val := range md.Get(cookieHeader) {
		parts := strings.Split(val, headerSep)
		if len(parts) != headerParts {
			continue
		}

		if parts[0] == cookieName {
			return parts[1]
		}
	}

	return ""
}

func (i *Implementation) readHTTPCookie(r *http.Request) string {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return ""
	}

	return cookie.Value
}

func (i *Implementation) setGRPCCookie(ctx context.Context, value string, drop bool) error {
	cookie := i.cookie(value, drop)

	err := grpc.SetHeader(ctx, metadata.New(map[string]string{"Set-Cookie": cookie.String()}))
	if err != nil {
		return simplerr.Wrap(err, "can't set header")
	}

	return nil
}

func (i *Implementation) cookie(value string, drop bool) *http.Cookie {
	cookie := &http.Cookie{Name: cookieName, Value: value, HttpOnly: true, Path: "/", SameSite: http.SameSiteLaxMode}

	if drop {
		cookie.MaxAge = -1
	}

	return cookie
}
