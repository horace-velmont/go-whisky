package servicefx

import (
	whiskySvc "github.com/GagulProject/go-whisky/internal/service/whisky"
	"go.uber.org/fx"
)

var Option = fx.Options(
	fx.Provide(whiskySvc.NewWhiskyService),
)
