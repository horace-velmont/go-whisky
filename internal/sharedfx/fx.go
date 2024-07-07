package sharedfx

import (
	"github.com/GagulProject/go-whisky/internal/shared"
	"go.uber.org/fx"
)

var Option = fx.Options(
	fx.Provide(
		shared.NewPSQL,
	),
)
