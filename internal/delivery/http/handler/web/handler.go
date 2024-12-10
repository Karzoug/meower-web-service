package web

import (
	"net/http"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"

	gen "github.com/Karzoug/meower-web-service/internal/delivery/http/gen/web/v1"
	"github.com/Karzoug/meower-web-service/internal/delivery/http/handler/errfunc"
	zerologHook "github.com/Karzoug/meower-web-service/internal/delivery/http/handler/zerolog"
	"github.com/Karzoug/meower-web-service/internal/delivery/http/middleware"
	"github.com/Karzoug/meower-web-service/internal/usecase"
)

const baseURL = "/api/web/v1"

var _ gen.StrictServerInterface = handlers{}

func RoutesFunc(usersUsecase usecase.UsersUseCase, postsUsecase usecase.PostsUseCase, tracer trace.Tracer, logger zerolog.Logger) func(mux *http.ServeMux) {
	logger = logger.With().
		Str("component", "http server: post handlers").
		Logger().
		Hook(zerologHook.TraceIDHook())

	hdl := handlers{
		usersUsecase: usersUsecase,
		postsUsecase: postsUsecase,
		logger:       logger,
	}

	return func(mux *http.ServeMux) {
		gen.HandlerWithOptions(
			gen.NewStrictHandlerWithOptions(hdl,
				[]gen.StrictMiddlewareFunc{
					middleware.Recover,
					middleware.AuthN,
					middleware.Error(logger),
					middleware.Logger(logger),
					middleware.Otel(tracer),
				},
				gen.StrictHTTPServerOptions{
					RequestErrorHandlerFunc:  errfunc.JSONRequest(logger, tracer),
					ResponseErrorHandlerFunc: errfunc.JSONResponse(logger),
				}),
			gen.StdHTTPServerOptions{
				BaseURL:    baseURL,
				BaseRouter: mux,
			})
	}
}

type handlers struct {
	postsUsecase usecase.PostsUseCase
	usersUsecase usecase.UsersUseCase
	logger       zerolog.Logger
}
