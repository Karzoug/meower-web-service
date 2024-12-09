package config

import (
	"github.com/rs/zerolog"

	"github.com/Karzoug/meower-common-go/metric/prom"
	"github.com/Karzoug/meower-common-go/trace/otlp"
)

type Config struct {
	LogLevel zerolog.Level     `env:"LOG_LEVEL" envDefault:"info"`
	PromHTTP prom.ServerConfig `envPrefix:"PROM_"`
	OTLP     otlp.Config       `envPrefix:"OTLP_"`
}
