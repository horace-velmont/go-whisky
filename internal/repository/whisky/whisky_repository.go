package whisky

import (
	"context"
	"database/sql"
	"github.com/GagulProject/go-whisky/generated/models"
	"github.com/GagulProject/go-whisky/internal/model/whisky"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type WhiskyRepository interface {
	Create(context.Context, *whisky.Whisky) (*models.Whisky, error)
}

type whiskyRepository struct {
	db *sql.DB
}

func NewWhiskyRepository(db *sql.DB) WhiskyRepository {
	return &whiskyRepository{
		db: db,
	}
}

func (r whiskyRepository) Create(ctx context.Context, whisky *whisky.Whisky) (*models.Whisky, error) {
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
