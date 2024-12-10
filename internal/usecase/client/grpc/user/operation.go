package user

import (
	"context"
	"log"

	"github.com/rs/xid"

	"github.com/Karzoug/meower-web-service/internal/entity"
	"github.com/Karzoug/meower-web-service/internal/usecase/client/grpc"
	userApi "github.com/Karzoug/meower-web-service/pkg/proto/grpc/user/v1"
)

func (c Client) GetProjection(ctx context.Context, authUserID xid.ID, userID string) (entity.UserProjection, error) {
	ctx = grpc.ContextWithUserID(ctx, authUserID)

	resp, err := c.c.GetShortProjection(ctx, &userApi.GetShortProjectionRequest{
		ByOneof: &userApi.GetShortProjectionRequest_Id{
			Id: userID,
		},
	})
	if err != nil {
		return entity.UserProjection{}, err
	}

	return fromProtoUser(resp), nil
}

func (c Client) BatchGetProjections(ctx context.Context, authUserID xid.ID, userIDs []string) ([]entity.UserProjection, error) {
	ctx = grpc.ContextWithUserID(ctx, authUserID)

	resp, err := c.c.BatchGetShortProjections(ctx, &userApi.BatchGetShortProjectionsRequest{
		Ids: userIDs,
	})
	if err != nil {
		return nil, err
	}

	log.Printf("got %d users", len(resp.Users))

	return fromProtoUsers(resp.Users), nil
}

func fromProtoUsers(u []*userApi.UserShortProjection) []entity.UserProjection {
	users := make([]entity.UserProjection, len(u))
	for i := range u {
		users[i] = fromProtoUser(u[i])
	}
	return users
}

func fromProtoUser(u *userApi.UserShortProjection) entity.UserProjection {
	return entity.UserProjection{
		ID:         u.Id,
		Username:   u.Username,
		Name:       u.Name,
		ImageUrl:   u.ImageUrl,
		StatusText: u.StatusText,
	}
}
