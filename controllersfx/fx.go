package controllersfx

import (
	"github.com/GagulProject/go-whisky/controllers/whisky"
	"github.com/GagulProject/go-whisky/internal/http"
	"go.uber.org/fx"
)

var Option = fx.Options(
	fx.Provide(
		fx.Annotate(
			whisky.NewWhiskyController,
			fx.As(new(http.Routes)),
			fx.ResultTags(`group:"public.http.routes"`),
		),
	),
)
