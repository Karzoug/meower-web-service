package web

import (
	"context"
	"slices"

	"github.com/Karzoug/meower-common-go/auth"

	gen "github.com/Karzoug/meower-web-service/internal/delivery/http/gen/web/v1"
	"github.com/Karzoug/meower-web-service/internal/usecase/option"
)

// Returns a collection of recent posts by the user and users they follow.
// (GET /users/{id}/timeline).
func (h handlers) GetUsersIdTimeline(ctx context.Context, request gen.GetUsersIdTimelineRequestObject) (gen.GetUsersIdTimelineResponseObject, error) {
	pgn := option.Pagination{}
	if request.Params.MaxPageSize != nil {
		pgn.Size = *request.Params.MaxPageSize
	}
	if request.Params.PageToken != nil {
		pgn.Token = *request.Params.PageToken
	}

	opts := option.ReturnPost{}
	if request.Params.Expansions != nil {
		if slices.Contains(*request.Params.Expansions, gen.AuthorId) {
			opts.IncludeUser = true
		}
	}

	res, err := h.postsUsecase.ListHomeTimeline(ctx, auth.UserIDFromContext(ctx), opts, pgn)
	if err != nil {
		return nil, err
	}

	resp := gen.GetUsersIdTimeline200JSONResponse{
		Data: toGenPosts(res.Posts),
	}

	if len(res.Authors) != 0 {
		incl := struct {
			Users []gen.UserShortProjection `json:"users,omitempty"`
		}{
			Users: toGenUserShortProjections(res.Authors),
		}
		resp.Includes = &incl
	}

	return resp, nil
}

// Returns posts composed by a single user, specified by the requested user ID.
// (GET /users/{id}/posts).
func (h handlers) GetUsersIdPosts(ctx context.Context, request gen.GetUsersIdPostsRequestObject) (gen.GetUsersIdPostsResponseObject, error) {
	pgn := option.Pagination{}
	if request.Params.MaxPageSize != nil {
		pgn.Size = *request.Params.MaxPageSize
	}
	if request.Params.PageToken != nil {
		pgn.Token = *request.Params.PageToken
	}

	opts := option.ReturnPost{}
	if request.Params.Expansions != nil {
		if slices.Contains(*request.Params.Expansions, gen.AuthorId) {
			opts.IncludeUser = true
		}
	}

	res, err := h.postsUsecase.ListUserTimeline(ctx, auth.UserIDFromContext(ctx), request.Id, opts, pgn)
	if err != nil {
		return nil, err
	}

	resp := gen.GetUsersIdPosts200JSONResponse{
		Data: toGenPosts(res.Posts),
	}

	if len(res.Authors) != 0 {
		incl := struct {
			Users []gen.UserShortProjection `json:"users,omitempty"`
		}{
			Users: toGenUserShortProjections(res.Authors),
		}
		resp.Includes = &incl
	}

	return resp, nil
}
