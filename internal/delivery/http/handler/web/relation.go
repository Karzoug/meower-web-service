package web

import (
	"context"
	"net/http"

	gen "github.com/Karzoug/meower-web-service/internal/delivery/http/gen/web/v1"
	"github.com/Karzoug/meower-web-service/internal/delivery/http/httperr"
)

// Allows a user ID to unfollow another user.
// (DELETE /users/{sourceUserId}/following/{targetUserId}).
func (h handlers) DeleteUsersSourceUserIdFollowingTargetUserId(ctx context.Context, request gen.DeleteUsersSourceUserIdFollowingTargetUserIdRequestObject) (gen.DeleteUsersSourceUserIdFollowingTargetUserIdResponseObject, error) {
	return nil, httperr.NewError("not implemented", http.StatusNotImplemented)
}

// Allows an authenticated user ID to unmute the target user.
// (DELETE /users/{sourceUserId}/muting/{targetUserId}).
func (h handlers) DeleteUsersSourceUserIdMutingTargetUserId(ctx context.Context, request gen.DeleteUsersSourceUserIdMutingTargetUserIdRequestObject) (gen.DeleteUsersSourceUserIdMutingTargetUserIdResponseObject, error) {
	return nil, httperr.NewError("not implemented", http.StatusNotImplemented)
}

// Allows a user ID to follow another user.
// (POST /users/{id}/following).
func (h handlers) PostUsersIdFollowing(ctx context.Context, request gen.PostUsersIdFollowingRequestObject) (gen.PostUsersIdFollowingResponseObject, error) {
	return nil, httperr.NewError("not implemented", http.StatusNotImplemented)
}

// Allows an authenticated user ID to mute the target user.
// (POST /users/{id}/muting).
func (h handlers) PostUsersIdMuting(ctx context.Context, request gen.PostUsersIdMutingRequestObject) (gen.PostUsersIdMutingResponseObject, error) {
	return nil, httperr.NewError("not implemented", http.StatusNotImplemented)
}
