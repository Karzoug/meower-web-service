package web

import (
	"context"
	"net/http"

	"github.com/Karzoug/meower-common-go/auth"
	"github.com/rs/xid"

	gen "github.com/Karzoug/meower-web-service/internal/delivery/http/gen/web/v1"
	"github.com/Karzoug/meower-web-service/internal/delivery/http/httperr"
	"github.com/Karzoug/meower-web-service/internal/entity"
)

// Returns a variety of information about user specified by username.
// (GET /users/by/username/{username}).
func (h handlers) GetUsersByUsernameUsername(ctx context.Context, request gen.GetUsersByUsernameUsernameRequestObject) (gen.GetUsersByUsernameUsernameResponseObject, error) {
	u, err := h.usersUsecase.GetUserWithUsername(ctx, request.Username)
	if err != nil {
		return nil, err
	}

	return gen.GetUsersByUsernameUsername200JSONResponse(toGenUserShortProjection(u)), nil
}

// Returns information about an authorized user.
// (GET /users/me).
func (h handlers) GetUsersMe(ctx context.Context, request gen.GetUsersMeRequestObject) (gen.GetUsersMeResponseObject, error) {
	u, err := h.usersUsecase.GetMe(ctx, auth.UserIDFromContext(ctx))
	if err != nil {
		return nil, err
	}

	return gen.GetUsersMe200JSONResponse(toGenUser(u)), nil
}

// Update information about an authorized user.
// (PUT /users/me).
func (h handlers) PutUsersMe(ctx context.Context, request gen.PutUsersMeRequestObject) (gen.PutUsersMeResponseObject, error) {
	if request.Body == nil {
		return nil, httperr.NewError("request body is nil", http.StatusBadRequest)
	}

	u, err := fromGenUser(*request.Body)
	if err != nil {
		return nil, httperr.NewError("invalid request body", http.StatusBadRequest)
	}

	if err := h.usersUsecase.UpdateMe(ctx, auth.UserIDFromContext(ctx), u); err != nil {
		return nil, err
	}

	return gen.PutUsersMe200Response{}, nil
}

// Returns a variety of information about a single user specified by the requested ID.
// (GET /users/{id}).
func (h handlers) GetUsersId(ctx context.Context, request gen.GetUsersIdRequestObject) (gen.GetUsersIdResponseObject, error) {
	userID, err := xid.FromString(request.Id)
	if err != nil {
		return nil, httperr.NewError("invalid user id format", http.StatusBadRequest)
	}

	u, err := h.usersUsecase.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return gen.GetUsersId200JSONResponse(toGenUserShortProjection(u)), nil
}

func toGenUserShortProjections(u []entity.UserProjection) []gen.UserShortProjection {
	users := make([]gen.UserShortProjection, len(u))
	for i := range u {
		users[i] = toGenUserShortProjection(u[i])
	}
	return users
}

func toGenUserShortProjection(u entity.UserProjection) gen.UserShortProjection {
	return gen.UserShortProjection{
		Username:   u.Username,
		Name:       u.Name,
		ImageUrl:   notEmptyOrNil(u.ImageUrl),
		StatusText: notEmptyOrNil(u.StatusText),
		Id:         u.ID,
	}
}

func toGenUser(u entity.User) gen.User {
	return gen.User{
		Username:   u.Username,
		Name:       u.Name,
		ImageUrl:   u.ImageUrl,
		StatusText: u.StatusText,
		Id:         u.ID,
	}
}

func fromGenUser(u gen.User) (entity.User, error) {
	return entity.User{
		ID:         u.Id,
		Username:   u.Username,
		Name:       u.Name,
		ImageUrl:   u.ImageUrl,
		StatusText: u.StatusText,
	}, nil
}

func notEmptyOrNil[T comparable](v T) *T {
	var t T
	if v == t {
		return nil
	}
	return &v
}
