package config

import (
	"github.com/rs/zerolog"

	"github.com/Karzoug/meower-common-go/metric/prom"
	"github.com/Karzoug/meower-common-go/trace/otlp"

	httpConfig "github.com/Karzoug/meower-web-service/internal/delivery/http/config"
)

type Config struct {
	LogLevel zerolog.Level           `env:"LOG_LEVEL" envDefault:"info"`
	HTTP     httpConfig.ServerConfig `envPrefix:"HTTP_"`
	PromHTTP prom.ServerConfig       `envPrefix:"PROM_"`
	OTLP     otlp.Config             `envPrefix:"OTLP_"`
}
