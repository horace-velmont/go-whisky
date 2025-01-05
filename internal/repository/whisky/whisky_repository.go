package whisky

import (
	"context"
	"database/sql"
	"github.com/GagulProject/go-whisky/generated/models"
	"github.com/GagulProject/go-whisky/internal/model/whisky"
	"github.com/GagulProject/go-whisky/internal/shared/errors"
	"github.com/GagulProject/go-whisky/internal/shared/scroller"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"time"
)

type WhiskyRepository interface {
	Create(context.Context, *whisky.Whisky) (*models.Whisky, error)
	ScrollAll(context.Context, ...scroller.ScrollOptionFn[time.Time]) (*scroller.Page[*whisky.Whisky], error)
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

func (r whiskyRepository) ScrollAll(ctx context.Context, optionFns ...scroller.ScrollOptionFn[time.Time]) (*scroller.Page[*whisky.Whisky], error) {
	option := scroller.NewScrollOption(optionFns...)

	page, err := scroller.Scroll[*whisky.Whisky](
		ctx,
		r.db,
		option,
		models.WhiskyColumns.CreatedAt,
	)
	return page, errors.Wrap(err)
}

func toBoiler(whisky *whisky.Whisky) *models.Whisky {
	return &models.Whisky{
		Strength:  whisky.Strength,
		Size:      whisky.Size,
		CreatedAt: whisky.CreatedAt,
		UpdatedAt: whisky.CreatedAt,
	}
}
