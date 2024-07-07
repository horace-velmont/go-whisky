package repositoryfx

import (
	"github.com/GagulProject/go-whisky/internal/repository/whisky"
	"go.uber.org/fx"
)

var Option = fx.Options(
	fx.Provide(whisky.NewWhiskyRepositories),
)
