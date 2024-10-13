package controllersfx

import (
	"github.com/GagulProject/go-whisky/controllers"
	"github.com/GagulProject/go-whisky/internal/http"
	"go.uber.org/fx"
)

var Option = fx.Options(
	fx.Provide(
		fx.Annotate(
			controllers.NewWhiskyController,
			fx.As(new(http.Routes)),
			fx.ResultTags(`group:"public.http.routes"`),
		),
	),
)
