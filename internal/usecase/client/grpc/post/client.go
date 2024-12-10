package post

import (
	"fmt"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gcfg "github.com/Karzoug/meower-web-service/internal/usecase/client/grpc"
	postApi "github.com/Karzoug/meower-web-service/pkg/proto/grpc/post/v1"
)

type Client struct {
	c postApi.PostServiceClient
}

func NewServiceClient(cfg gcfg.Config) (Client, error) {
	conn, err := grpc.NewClient(
		cfg.URI,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		return Client{}, fmt.Errorf("could not connect to post service: %w", err)
	}
	return Client{c: postApi.NewPostServiceClient(conn)}, nil
}
