package relation

import (
	"fmt"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gcfg "github.com/Karzoug/meower-web-service/internal/usecase/client/grpc"
	relationApi "github.com/Karzoug/meower-web-service/pkg/proto/grpc/relation/v1"
)

type Client struct {
	c relationApi.RelationServiceClient
}

func NewServiceClient(cfg gcfg.Config) (Client, error) {
	conn, err := grpc.NewClient(
		cfg.URI,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		return Client{}, fmt.Errorf("could not connect to relation service: %w", err)
	}
	return Client{c: relationApi.NewRelationServiceClient(conn)}, nil
}
