package mymiddleware

import (
	"context"
	"net/http"

	"app/app"
	"github.com/goadesign/goa"
)

var (
	// ErrUnauthorized is the error returned for unauthorized requests.
	ErrUnauthorized = goa.NewErrorClass("unauthorized", 401)
	// ErrInternalServer is the error returned for internal error.
	ErrInternalServer = goa.NewErrorClass("internal error", 500)
)

// NewAdminUserAuthMiddleware token check authority type is admin
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

// NewGeneralUserAuthMiddleware token check authority type is general
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

// NewTestModeMiddleware token check none executing
func NewTestModeMiddleware() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			return h(ctx, rw, req)
		}
	}
}
