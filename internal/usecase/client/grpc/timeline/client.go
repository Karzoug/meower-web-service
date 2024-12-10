package timeline

import (
	"fmt"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gcfg "github.com/Karzoug/meower-web-service/internal/usecase/client/grpc"
	timelineApi "github.com/Karzoug/meower-web-service/pkg/proto/grpc/timeline/v1"
)

type Client struct {
	c timelineApi.TimelineServiceClient
}

func NewServiceClient(cfg gcfg.Config) (Client, error) {
	conn, err := grpc.NewClient(
		cfg.URI,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		return Client{}, fmt.Errorf("could not connect to timeline service: %w", err)
	}
	return Client{c: timelineApi.NewTimelineServiceClient(conn)}, nil
}
