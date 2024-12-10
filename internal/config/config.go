package config

import (
	"github.com/rs/zerolog"

	"github.com/Karzoug/meower-common-go/metric/prom"
	"github.com/Karzoug/meower-common-go/trace/otlp"

	httpConfig "github.com/Karzoug/meower-web-service/internal/delivery/http/config"
	"github.com/Karzoug/meower-web-service/internal/usecase/client/grpc"
)

type Config struct {
	LogLevel        zerolog.Level           `env:"LOG_LEVEL" envDefault:"info"`
	HTTP            httpConfig.ServerConfig `envPrefix:"HTTP_"`
	PromHTTP        prom.ServerConfig       `envPrefix:"PROM_"`
	OTLP            otlp.Config             `envPrefix:"OTLP_"`
	PostService     grpc.Config             `envPrefix:"POST_SERVICE_"`
	TimelineService grpc.Config             `envPrefix:"TIMELINE_SERVICE_"`
	UserService     grpc.Config             `envPrefix:"USER_SERVICE_"`
}
