package app

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel"
	"golang.org/x/sync/errgroup"

	"github.com/Karzoug/meower-common-go/metric/prom"
	"github.com/Karzoug/meower-common-go/trace/otlp"

	"github.com/Karzoug/meower-web-service/internal/config"
	healthHandlers "github.com/Karzoug/meower-web-service/internal/delivery/http/handler/health"
	webHttp "github.com/Karzoug/meower-web-service/internal/delivery/http/handler/web"
	httpServer "github.com/Karzoug/meower-web-service/internal/delivery/http/server"
	"github.com/Karzoug/meower-web-service/internal/usecase"
	"github.com/Karzoug/meower-web-service/internal/usecase/client/grpc/post"
	"github.com/Karzoug/meower-web-service/internal/usecase/client/grpc/timeline"
	"github.com/Karzoug/meower-web-service/internal/usecase/client/grpc/user"
	"github.com/Karzoug/meower-web-service/pkg/buildinfo"
)

const (
	serviceName     = "WebService"
	metricNamespace = "web_service"
	pkgName         = "github.com/Karzoug/meower-web-service"
	initTimeout     = 10 * time.Second
	shutdownTimeout = 10 * time.Second
)

var serviceVersion = buildinfo.Get().ServiceVersion

func Run(ctx context.Context, logger zerolog.Logger) error {
	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(cfg.LogLevel)

	logger.Info().
		Int("GOMAXPROCS", runtime.GOMAXPROCS(0)).
		Str("log level", cfg.LogLevel.String()).
		Msg("starting up")

	// set timeout for initialization
	ctxInit, closeCtx := context.WithTimeout(ctx, initTimeout)
	defer closeCtx()

	// set up tracer
	cfg.OTLP.ServiceName = serviceName
	cfg.OTLP.ServiceVersion = serviceVersion
	cfg.OTLP.ExcludedHTTPRoutes = map[string]struct{}{
		"/readiness": {},
		"/liveness":  {},
	}
	shutdownTracer, err := otlp.RegisterGlobal(ctxInit, cfg.OTLP)
	if err != nil {
		return err
	}
	defer doClose(shutdownTracer, logger)

	tracer := otel.GetTracerProvider().Tracer(pkgName)

	// set up meter
	shutdownMeter, err := prom.RegisterGlobal(ctxInit, serviceName, serviceVersion, metricNamespace)
	if err != nil {
		return err
	}
	defer doClose(shutdownMeter, logger)

	// set up user service grpc client
	userClient, err := user.NewServiceClient(cfg.UserService)
	if err != nil {
		return fmt.Errorf("could not connect to post microservice: %w", err)
	}

	// set up post service grpc client
	postClient, err := post.NewServiceClient(cfg.PostService)
	if err != nil {
		return fmt.Errorf("could not connect to post microservice: %w", err)
	}

	// set up timeline service grpc client
	timelineClient, err := timeline.NewServiceClient(cfg.TimelineService)
	if err != nil {
		return fmt.Errorf("could not connect to relation microservice: %w", err)
	}

	usersUsecase := usecase.NewUsersUseCase(userClient)
	postUsecases := usecase.NewPostsUseCase(postClient, userClient, timelineClient, logger)

	// set up http server
	httpSrv := httpServer.New(
		cfg.HTTP,
		[]httpServer.Routes{
			webHttp.RoutesFunc(usersUsecase, postUsecases, tracer, logger),
			healthHandlers.RoutesFunc(logger),
		},
		logger)

	eg, ctx := errgroup.WithContext(ctx)
	// run service http server
	eg.Go(func() error {
		return httpSrv.Run(ctx)
	})
	// run prometheus metrics http server
	eg.Go(func() error {
		return prom.Serve(ctx, cfg.PromHTTP, logger)
	})

	return eg.Wait()
}

func doClose(fn func(context.Context) error, logger zerolog.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := fn(ctx); err != nil {
		logger.Error().
			Err(err).
			Msg("error closing")
	}
}
