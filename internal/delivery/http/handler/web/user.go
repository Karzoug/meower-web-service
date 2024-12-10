package web

import (
	"context"
	"net/http"

	gen "github.com/Karzoug/meower-web-service/internal/delivery/http/gen/web/v1"
	"github.com/Karzoug/meower-web-service/internal/delivery/http/httperr"
)

// Returns a variety of information about user specified by username.
// (GET /users/by/username/{username}).
func (h handlers) GetUsersByUsernameUsername(ctx context.Context, request gen.GetUsersByUsernameUsernameRequestObject) (gen.GetUsersByUsernameUsernameResponseObject, error) {
	return nil, httperr.NewError("not implemented", http.StatusNotImplemented)
}

// Returns information about an authorized user.
// (GET /users/me).
func (h handlers) GetUsersMe(ctx context.Context, request gen.GetUsersMeRequestObject) (gen.GetUsersMeResponseObject, error) {
	return nil, httperr.NewError("not implemented", http.StatusNotImplemented)
}

// Update information about an authorized user.
// (PUT /users/me).
func (h handlers) PutUsersMe(ctx context.Context, request gen.PutUsersMeRequestObject) (gen.PutUsersMeResponseObject, error) {
	return nil, httperr.NewError("not implemented", http.StatusNotImplemented)
}

// Returns a variety of information about a single user specified by the requested ID.
// (GET /users/{id}).
func (h handlers) GetUsersId(ctx context.Context, request gen.GetUsersIdRequestObject) (gen.GetUsersIdResponseObject, error) {
	return nil, httperr.NewError("not implemented", http.StatusNotImplemented)
}
