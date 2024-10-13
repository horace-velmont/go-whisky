package main

import (
	"github.com/GagulProject/go-whisky/internal/http"
	"github.com/GagulProject/go-whisky/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		server.HTTPServerModule(),
		fx.Module(
			"internal",
			server.Options(),
		),
		fx.Invoke(
			fx.Annotate(
				func(_ []*http.Server) error {
					return nil
				},
				fx.ParamTags(`group:"http.servers"`),
			),
		),
	).Run()
}
