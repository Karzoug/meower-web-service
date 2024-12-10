package web

import (
	"context"
	"net/http"
	"slices"

	"github.com/Karzoug/meower-common-go/auth"
	"github.com/rs/xid"

	gen "github.com/Karzoug/meower-web-service/internal/delivery/http/gen/web/v1"
	"github.com/Karzoug/meower-web-service/internal/delivery/http/httperr"
	"github.com/Karzoug/meower-web-service/internal/usecase"
)

// Returns a collection of recent posts by the user and users they follow.
// (GET /users/{id}/timeline).
func (h handlers) GetUsersIdTimeline(ctx context.Context, request gen.GetUsersIdTimelineRequestObject) (gen.GetUsersIdTimelineResponseObject, error) {
	const op = "GET /users/{id}/timeline"

	userID := auth.UserIDFromContext(ctx)
	if userID.IsZero() {
		h.logger.Error().
			Str("operation", op).
			Msg("bug: user id is nil")
		return nil, httperr.NewError("empty auth data", http.StatusUnauthorized)
	}

	pgn := usecase.PaginationOptions{}
	if request.Params.MaxPageSize != nil {
		pgn.Size = *request.Params.MaxPageSize
	}
	if request.Params.PageToken != nil {
		pgn.Token = *request.Params.PageToken
	}

	opts := usecase.ReturnPostOptions{}
	if request.Params.Expansions != nil {
		if slices.Contains(*request.Params.Expansions, gen.AuthorId) {
			opts.IncludeUser = true
		}
	}

	posts, err := h.postsUsecase.ListHomeTimeline(ctx, userID, opts, pgn)
	if err != nil {
		return nil, err
	}

	return gen.GetUsersIdTimeline200JSONResponse{
		Data: toGenPosts(posts),
	}, nil
}

// Returns posts composed by a single user, specified by the requested user ID.
// (GET /users/{id}/posts).
func (h handlers) GetUsersIdPosts(ctx context.Context, request gen.GetUsersIdPostsRequestObject) (gen.GetUsersIdPostsResponseObject, error) {
	userID, err := xid.FromString(request.Id)
	if err != nil {
		return nil, httperr.NewError("invalid user id format", http.StatusBadRequest)
	}

	pgn := usecase.PaginationOptions{}
	if request.Params.MaxPageSize != nil {
		pgn.Size = *request.Params.MaxPageSize
	}
	if request.Params.PageToken != nil {
		pgn.Token = *request.Params.PageToken
	}

	opts := usecase.ReturnPostOptions{}
	if request.Params.Expansions != nil {
		if slices.Contains(*request.Params.Expansions, gen.AuthorId) {
			opts.IncludeUser = true
		}
	}

	posts, err := h.postsUsecase.ListUserTimeline(ctx, userID, opts, pgn)
	if err != nil {
		return nil, err
	}

	return gen.GetUsersIdPosts200JSONResponse{
		Data: toGenPosts(posts),
	}, nil
}
