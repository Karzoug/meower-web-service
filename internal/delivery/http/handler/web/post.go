package web

import (
	"context"
	"net/http"
	"slices"

	"github.com/Karzoug/meower-common-go/auth"
	"github.com/rs/xid"

	gen "github.com/Karzoug/meower-web-service/internal/delivery/http/gen/web/v1"
	"github.com/Karzoug/meower-web-service/internal/delivery/http/httperr"
	"github.com/Karzoug/meower-web-service/internal/entity"
	"github.com/Karzoug/meower-web-service/internal/usecase"
)

// Creates a post on behalf of an authenticated user.
// (POST /posts/).
func (h handlers) PostPosts(ctx context.Context, request gen.PostPostsRequestObject) (gen.PostPostsResponseObject, error) {
	const op = "POST /posts/"

	if request.Body == nil {
		return nil, httperr.NewError("request body is nil", http.StatusBadRequest)
	}

	userID := auth.UserIDFromContext(ctx)
	if userID.IsZero() {
		h.logger.Error().
			Str("operation", op).
			Msg("bug: user id is nil")
		return nil, httperr.NewError("empty auth data", http.StatusUnauthorized)
	}

	p, err := h.postsUsecase.CreatePost(ctx, userID, entity.NewPost{
		Text: request.Body.Text,
	})
	if err != nil {
		return nil, err
	}

	return gen.PostPosts201JSONResponse{
		Id: p.ID.String(),
	}, nil
}

// Delete a post by the requested ID.
// (DELETE /posts/{id}).
func (h handlers) DeletePostsId(ctx context.Context, request gen.DeletePostsIdRequestObject) (gen.DeletePostsIdResponseObject, error) {
	const op = "DELETE /posts/{id}"

	postID, err := xid.FromString(request.Id)
	if err != nil {
		return nil, httperr.NewError("invalid post id format", http.StatusBadRequest)
	}

	userID := auth.UserIDFromContext(ctx)
	if userID.IsZero() {
		h.logger.Error().
			Str("operation", op).
			Msg("bug: user id is nil")
		return nil, httperr.NewError("empty auth data", http.StatusUnauthorized)
	}

	if err := h.postsUsecase.DeletePost(ctx, userID, postID); err != nil {
		return nil, err
	}

	return gen.DeletePostsId200Response{}, nil
}

// Returns a variety of information about a single post specified by the requested ID.
// (GET /posts/{id}).
func (h handlers) GetPostsId(ctx context.Context, request gen.GetPostsIdRequestObject) (gen.GetPostsIdResponseObject, error) {
	// const op = "GET /posts/{id}"

	h.logger.Debug().
		Any("request", request).
		Msg("got request")

	postID, err := xid.FromString(request.Id)
	if err != nil {
		return nil, httperr.NewError("invalid post id format", http.StatusBadRequest)
	}

	opts := usecase.ReturnPostOptions{}
	if request.Params.Expansions != nil {
		if slices.Contains(*request.Params.Expansions, gen.AuthorId) {
			opts.IncludeUser = true
		}
	}

	p, err := h.postsUsecase.GetPost(ctx, postID, opts)
	if err != nil {
		return nil, err
	}

	return gen.GetPostsId200JSONResponse{
		Data: toGenPost(p),
	}, nil
}

func toGenPost(post entity.Post) gen.Post {
	return gen.Post{
		AuthorId:  post.ID.String(),
		Id:        post.ID.String(),
		Text:      post.Text,
		UpdatedAt: post.UpdatedAt,
	}
}

func toGenPosts(posts []entity.Post) []gen.Post {
	res := make([]gen.Post, len(posts))
	for i, p := range posts {
		res[i] = toGenPost(p)
	}
	return res
}
