package grpc

type Config struct {
	URI string `env:"URI,notEmpty"`
}
