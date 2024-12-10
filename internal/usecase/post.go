package usecase

import (
	"context"

	"github.com/Karzoug/meower-common-go/ucerr"
	"github.com/Karzoug/meower-web-service/internal/entity"
	"github.com/rs/xid"
	"google.golang.org/grpc/codes"
)

type PostsUseCase struct {
}

func NewPostsUseCase() PostsUseCase {
	return PostsUseCase{}
}

// Creates a post on behalf of an authenticated user.
func (uc PostsUseCase) CreatePost(ctx context.Context, authUserID xid.ID, post entity.NewPost) (entity.Post, error) {
	return entity.Post{}, ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}

// Delete a post on behalf of an authenticated user ID.
func (uc PostsUseCase) DeletePost(ctx context.Context, authUserID, postID xid.ID) error {
	return ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}

func (uc PostsUseCase) GetPost(ctx context.Context, postID xid.ID, opts ReturnPostOptions) (entity.Post, error) {

	return entity.Post{}, ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}

func (uc PostsUseCase) ListUserTimeline(ctx context.Context, userID xid.ID, opts ReturnPostOptions, pgn PaginationOptions) ([]entity.Post, error) {
	return nil, ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}

func (uc PostsUseCase) ListHomeTimeline(ctx context.Context, userID xid.ID, opts ReturnPostOptions, pgn PaginationOptions) ([]entity.Post, error) {
	return nil, ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}
