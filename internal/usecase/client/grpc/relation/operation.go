package relation

import (
	"context"

	"github.com/rs/xid"

	"github.com/Karzoug/meower-web-service/internal/usecase/client/grpc"
	relationApi "github.com/Karzoug/meower-web-service/pkg/proto/grpc/relation/v1"
)

func (c Client) ListFollowerIDs(ctx context.Context, authUserID xid.ID, userID xid.ID) ([]xid.ID, error) {
	ctx = grpc.ContextWithUserID(ctx, authUserID)

	fresp, err := c.c.ListFollowers(ctx,
		&relationApi.ListFollowersRequest{
			Parent:   userID.String(),
			PageSize: -1, // require all followings
		})
	if err != nil {
		return nil, err
	}

	followerIDs := make([]xid.ID, 0, len(fresp.Followers))
	for i := range fresp.Followers {
		if fresp.Followers[i] == nil || fresp.Followers[i].Muted {
			continue
		}
		id, _ := xid.FromString(fresp.Followers[i].Id)
		followerIDs = append(followerIDs, id)
	}

	return followerIDs, nil
}

func (c Client) ListNotMutedFollowingIDs(ctx context.Context, authUserID xid.ID, userID xid.ID) ([]xid.ID, error) {
	ctx = grpc.ContextWithUserID(ctx, authUserID)

	fresp, err := c.c.ListFollowings(ctx,
		&relationApi.ListFollowingsRequest{
			Parent:   userID.String(),
			PageSize: -1, // require all followings
		})
	if err != nil {
		return nil, err
	}

	followingIDs := make([]xid.ID, 0, len(fresp.Followings))
	for i := range fresp.Followings {
		if fresp.Followings[i] == nil ||
			fresp.Followings[i].Muted {
			continue
		}

		id, _ := xid.FromString(fresp.Followings[i].Id)
		followingIDs = append(followingIDs, id)
	}

	return followingIDs, nil
}
