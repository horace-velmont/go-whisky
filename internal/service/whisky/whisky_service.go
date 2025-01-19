package whisky

import (
	"context"
	"github.com/GagulProject/go-whisky/internal/model/whisky"
	whiskyR "github.com/GagulProject/go-whisky/internal/repository/whisky"
	"github.com/GagulProject/go-whisky/internal/shared/epoch"
	"github.com/GagulProject/go-whisky/internal/shared/scroller"
	"time"
)

type whiskyService struct {
	repo whiskyR.WhiskyRepository
}

type WhiskyService interface {
	Create(context.Context, *whisky.Whisky) (*whisky.Whisky, error)
	ScrollAll(ctx context.Context, request *scroller.PageRequest[epoch.Milli]) (*scroller.Page[*whisky.Whisky], error)
}

func NewWhiskyService(repo whiskyR.WhiskyRepository) WhiskyService {
	return &whiskyService{
		repo: repo,
	}
}

func (w *whiskyService) Create(ctx context.Context, whisky *whisky.Whisky) (*whisky.Whisky, error) {
	return w.repo.Create(ctx, whisky)
}

func (w *whiskyService) ScrollAll(ctx context.Context, request *scroller.PageRequest[epoch.Milli]) (*scroller.Page[*whisky.Whisky], error) {
	return w.repo.ScrollAll(
		ctx,
		scroller.WithLimit[time.Time](request.Limit),
		scroller.WithOrder[time.Time](request.Order),
		scroller.WithSince(request.Since.ToTime()),
	)
}
