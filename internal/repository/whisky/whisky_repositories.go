package whisky

import (
	"context"
	"database/sql"
	"github.com/GagulProject/go-whisky/generated/models"
	"github.com/GagulProject/go-whisky/internal/model/whisky"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type WhiskyRepositories interface {
	Create(context.Context, *whisky.Whisky) (*models.Whisky, error)
}

type whiskyRepositories struct {
	db *sql.DB
}

func NewWhiskyRepositories(db *sql.DB) WhiskyRepositories {
	return &whiskyRepositories{
		db: db,
	}
}

func (r whiskyRepositories) Create(ctx context.Context, whisky *whisky.Whisky) (*models.Whisky, error) {
	model := toBoiler(whisky)
	err := model.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	err = model.Reload(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func toBoiler(whisky *whisky.Whisky) *models.Whisky {
	return &models.Whisky{
		Strength:  whisky.Strength,
		Size:      whisky.Size,
		CreatedAt: whisky.CreatedAt,
		UpdatedAt: whisky.CreatedAt,
	}
}
