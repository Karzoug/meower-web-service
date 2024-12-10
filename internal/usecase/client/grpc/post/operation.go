package post

import (
	"context"

	"github.com/rs/xid"

	"github.com/Karzoug/meower-web-service/internal/entity"
	"github.com/Karzoug/meower-web-service/internal/usecase/client/grpc"
	"github.com/Karzoug/meower-web-service/internal/usecase/option"
	postApi "github.com/Karzoug/meower-web-service/pkg/proto/grpc/post/v1"
)

func (c Client) Create(ctx context.Context, authUserID xid.ID, post entity.NewPost) (entity.Post, error) {
	ctx = grpc.ContextWithUserID(ctx, authUserID)

	res, err := c.c.CreatePost(ctx, &postApi.CreatePostRequest{Post: toProtoPost(post)})
	if err != nil {
		return entity.Post{}, err
	}

	return fromProtoPost(res), nil
}

func (c Client) Get(ctx context.Context, authUserID xid.ID, id string) (entity.Post, error) {
	ctx = grpc.ContextWithUserID(ctx, authUserID)

	res, err := c.c.GetPost(ctx, &postApi.GetPostRequest{Id: id})
	if err != nil {
		return entity.Post{}, err
	}

	return fromProtoPost(res), nil
}

func (c Client) List(ctx context.Context, authUserID xid.ID, userID string, pgn option.Pagination) ([]entity.Post, string, error) {
	ctx = grpc.ContextWithUserID(ctx, authUserID)

	req := &postApi.ListPostsRequest{
		Parent:    userID,
		PageToken: pgn.Token,
		PageSize:  int32(pgn.Size),
	}

	res, err := c.c.ListPosts(ctx, req)
	if err != nil {
		return nil, "", err
	}

	return fromProtoPosts(res.Posts), res.NextPageToken, nil
}

func toProtoPost(post entity.NewPost) *postApi.Post {
	return &postApi.Post{
		Text:     post.Text,
		AuthorId: post.AuthorID,
	}
}

func fromProtoPost(post *postApi.Post) entity.Post {
	p := entity.Post{
		ID:        post.Id,
		AuthorID:  post.AuthorId,
		UpdatedAt: post.UpdatedTime.AsTime(),
		IsDeleted: post.Deleted,
	}

	if p.IsDeleted {
		return p
	}

	p.Text = post.Text

	return p
}

func fromProtoPosts(posts []*postApi.Post) []entity.Post {
	res := make([]entity.Post, len(posts))
	for i, p := range posts {
		res[i] = fromProtoPost(p)
	}
	return res
}
