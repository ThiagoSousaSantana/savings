package main

import (
	"context"
	"net"
	"net/http"

	"github.com/ThiagoSousaSantana/saving/cmd/config"
	"github.com/ThiagoSousaSantana/saving/cmd/routes"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			config.NewConfig,
			NewHttpServer,
			fx.Annotate(
				NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
			AsRoute(routes.NewExpenseHandler),
			AsRoute(routes.NewIncomeHandler),
			zap.NewDevelopment,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

func NewHttpServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger, config *config.Config) *http.Server {
	srv := &http.Server{Addr: ":" + config.API.Port, Handler: mux}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Server started", zap.String("port", ln.Addr().String()))

			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Server stopped", zap.String("address", srv.Addr))
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

func NewServeMux(routes []routes.Route) *http.ServeMux {
	mux := http.NewServeMux()

	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}

	return mux
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(routes.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
