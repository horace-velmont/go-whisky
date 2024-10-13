package server

import (
	"context"
	"errors"
	"github.com/GagulProject/go-whisky/controllersfx"
	"github.com/GagulProject/go-whisky/internal/http"
	"github.com/GagulProject/go-whisky/internal/repositoryfx"
	"github.com/GagulProject/go-whisky/internal/sharedfx"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	netHttp "net/http"
	"time"
)

type Server struct {
	engine *gin.Engine
	server *http.Server
}

type ServerParams struct {
	fx.In

	Routes []http.Routes `group:"public.http.routes"`
}

type ServerResults struct {
	fx.Out

	Server *http.Server `group:"http.servers"`
}

func Options() fx.Option {
	return fx.Options(
		repositoryfx.Option,
		sharedfx.Option,
	)
}

func HTTPServerModule() fx.Option {
	return fx.Module(
		"httpServer",
		controllersfx.Option,

		fx.Provide(
			NewServer,
		),
	)
}

func NewServer(
	lifeCycle fx.Lifecycle,
	params ServerParams,
) (ServerResults, error) {
	server, err := http.NewServer(http.ServerConfig{Port: "8080"}, params.Routes)

	lifeCycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func(server *http.Server) {
				log.Println("starting server ...")
				if err := server.Run(); err != nil && !errors.Is(err, netHttp.ErrServerClosed) {
					panic(err)
				}
			}(server)
			return nil
		},
		OnStop: func(_ context.Context) error {
			log.Println("stopping server ...")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown:", err)
				return err
			}
			log.Println("stopped server success")
			return nil
		},
	})

	return ServerResults{
		Server: server,
	}, err
}
