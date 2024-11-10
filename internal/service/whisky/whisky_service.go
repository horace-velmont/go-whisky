package whisky

import (
	"context"
	"github.com/GagulProject/go-whisky/generated/models"
	"github.com/GagulProject/go-whisky/internal/model/whisky"
	whiskyR "github.com/GagulProject/go-whisky/internal/repository/whisky"
)

type whiskyService struct {
	repo whiskyR.WhiskyRepository
}

type WhiskyService interface {
	Create(context.Context, *whisky.Whisky) (*models.Whisky, error)
}

func NewWhiskyService(repo whiskyR.WhiskyRepository) WhiskyService {
	return &whiskyService{
		repo: repo,
	}
}

func (w *whiskyService) Create(ctx context.Context, whisky *whisky.Whisky) (*models.Whisky, error) {
	return w.repo.Create(ctx, whisky)
}
