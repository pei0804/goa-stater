package mymiddleware

import (
	"context"
	"net/http"

	"github.com/goadesign/goa"
	"github.com/pei0804/goa-stater/app"
)

var (
	// ErrUnauthorized is the error returned for unauthorized requests.
	ErrUnauthorized = goa.NewErrorClass("unauthorized", 401)
	// ErrUnauthorized is the error returned for unauthorized requests.
	ErrInternalServer = goa.NewErrorClass("internal error", 500)
)

func NewAdminUserAuthMiddleware() goa.Middleware {
	scheme := app.NewAdminAuthSecurity()
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log header specified by scheme
			token := req.Header.Get(scheme.Name)
			if len(token) == 0 {
				goa.LogError(ctx, "failed api token auth")
				return ErrUnauthorized("missing auth")
			}
			// Put your logic here

			goa.LogInfo(ctx, "auth", "apikey", "token", token)
			return h(ctx, rw, req)
		}
	}
}

func NewGeneralUserAuthMiddleware() goa.Middleware {
	scheme := app.NewGeneralAuthSecurity()
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log header specified by scheme
			token := req.Header.Get(scheme.Name)
			if len(token) == 0 {
				goa.LogError(ctx, "failed api token auth")
				return ErrUnauthorized("missing auth")
			}
			// Put your logic here

			goa.LogInfo(ctx, "auth", "apikey", "token", token)
			return h(ctx, rw, req)
		}
	}
}

// NewTestModeMiddleware テスト用のミドルウェア。テスト用の固定値をContextに格納する。
func NewTestModeMiddleware() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			return h(ctx, rw, req)
		}
	}
}
