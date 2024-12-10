package middleware

import (
	"context"
	"net/http"

	"github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	"github.com/rs/xid"

	"github.com/Karzoug/meower-common-go/auth"

	gen "github.com/Karzoug/meower-web-service/internal/delivery/http/gen/web/v1"
	"github.com/Karzoug/meower-web-service/internal/delivery/http/httperr"
)

const authHeader = "X-USER-ID"

// AuthN is a middleware that adds an username from the request "X-User-ID" Header to the context.
// (!) The middleware doesn't check if the token is valid - it's up to the outer gateway.
func AuthN(next nethttp.StrictHTTPHandlerFunc, operationID string) nethttp.StrictHTTPHandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request any) (resp any, err error) {
		sub := r.Header.Get(authHeader)
		if sub != "" {
			id, err := xid.FromString(sub)
			if err != nil {
				return nil, httperr.NewError("invalid user id format", http.StatusBadRequest)
			}
			ctx = auth.WithUserID(ctx, id)
		}

		// if spec claim authentification
		if ctx.Value(gen.OAuthScopes) != nil && sub == "" {
			return nil, httperr.NewError("unauthenticated", http.StatusUnauthorized)
		}

		return next(ctx, w, r, request)
	}
}
