package grpc

import (
	"context"

	"github.com/rs/xid"
	"google.golang.org/grpc/metadata"
)

const userKey string = "x-user-id"

func ContextWithUserID(ctx context.Context, userID xid.ID) context.Context {
	return metadata.AppendToOutgoingContext(ctx, userKey, userID.String())
}
