package usecase

import (
	"context"

	"github.com/Karzoug/meower-common-go/ucerr"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Karzoug/meower-web-service/internal/entity"
	"github.com/Karzoug/meower-web-service/internal/usecase/client/grpc/post"
	"github.com/Karzoug/meower-web-service/internal/usecase/client/grpc/timeline"
	"github.com/Karzoug/meower-web-service/internal/usecase/client/grpc/user"
	"github.com/Karzoug/meower-web-service/internal/usecase/option"
)

type PostsUseCase struct {
	postServiceClient     post.Client
	userServiceClient     user.Client
	timelineServiceClient timeline.Client
	logger                zerolog.Logger
}

func NewPostsUseCase(postClient post.Client,
	userServiceClient user.Client,
	timelineServiceClient timeline.Client,
	logger zerolog.Logger,
) PostsUseCase {
	return PostsUseCase{
		postServiceClient:     postClient,
		userServiceClient:     userServiceClient,
		timelineServiceClient: timelineServiceClient,
		logger:                logger,
	}
}

type PostGetResult struct {
	Post    entity.Post
	Authors []entity.UserProjection
}

type PostListResult struct {
	Posts   []entity.Post
	Authors []entity.UserProjection
	Token   string
}

// Creates a post on behalf of an authenticated user.
func (uc PostsUseCase) CreatePost(ctx context.Context, authUserID xid.ID, post entity.NewPost) (entity.Post, error) {
	p, err := uc.postServiceClient.Create(ctx, authUserID, post)
	if err != nil {
		st := status.Convert(err)
		return entity.Post{}, ucerr.NewError(err, st.Message(), st.Code())
	}

	return p, nil
}

// Delete a post on behalf of an authenticated user ID.
func (uc PostsUseCase) DeletePost(ctx context.Context, authUserID xid.ID, postID string) error {
	return ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}

func (uc PostsUseCase) GetPost(ctx context.Context, authUserID xid.ID, postID string, opts option.ReturnPost) (PostGetResult, error) {
	p, err := uc.postServiceClient.Get(ctx, authUserID, postID)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return PostGetResult{}, ucerr.NewError(err, st.Message(), st.Code())
		}
		return PostGetResult{}, ucerr.NewInternalError(err)
	}

	res := PostGetResult{Post: p}

	if opts.IncludeUser {
		u, err := uc.userServiceClient.GetProjection(ctx, authUserID, p.AuthorID)
		if err != nil {
			st := status.Convert(err)
			return PostGetResult{}, ucerr.NewError(err, st.Message(), st.Code())
		}

		res.Authors = []entity.UserProjection{u}
	}

	return res, nil
}

func (uc PostsUseCase) ListUserTimeline(ctx context.Context, authID xid.ID, userID string, opts option.ReturnPost, pgn option.Pagination) (PostListResult, error) {
	posts, token, err := uc.postServiceClient.List(ctx, authID, userID, pgn)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return PostListResult{}, ucerr.NewError(err, st.Message(), st.Code())
		}
		return PostListResult{}, ucerr.NewInternalError(err)
	}

	res := PostListResult{Posts: posts, Token: token}

	if opts.IncludeUser {
		ids := uniqueAuthorIDs(posts)
		uc.logger.Debug().Interface("ids", ids).Msg("author ids")
		users, err := uc.userServiceClient.BatchGetProjections(ctx, authID, ids)
		if err != nil {
			st := status.Convert(err)
			return PostListResult{}, ucerr.NewError(err, st.Message(), st.Code())
		}

		res.Authors = users
	}

	return res, nil
}

func (uc PostsUseCase) ListHomeTimeline(ctx context.Context, userID xid.ID, opts option.ReturnPost, pgn option.Pagination) (PostListResult, error) {
	return PostListResult{}, ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}

func uniqueAuthorIDs(posts []entity.Post) []string {
	userIDSet := make(map[string]struct{}, len(posts))
	for _, p := range posts {
		userIDSet[p.AuthorID] = struct{}{}
	}
	userIDs := make([]string, len(userIDSet))
	var i int
	for k := range userIDSet {
		userIDs[i] = k
		i++
	}

	return userIDs
}
