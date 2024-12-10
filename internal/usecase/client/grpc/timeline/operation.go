package timeline

import (
	"context"

	"github.com/rs/xid"

	"github.com/Karzoug/meower-web-service/internal/entity"
	"github.com/Karzoug/meower-web-service/internal/usecase/client/grpc"
)

func (c Client) ListHomeTimeline(ctx context.Context, authUserID xid.ID, userID string) ([]entity.TimelinePost, error) {
	ctx = grpc.ContextWithUserID(ctx, authUserID)

	return []entity.TimelinePost{}, nil
}
