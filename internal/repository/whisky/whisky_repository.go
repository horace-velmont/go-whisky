package whisky

import (
	"context"
	"database/sql"
	"time"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/GagulProject/go-whisky/generated/models"
	"github.com/GagulProject/go-whisky/internal/model/whisky"
	"github.com/GagulProject/go-whisky/internal/shared/errors"
	"github.com/GagulProject/go-whisky/internal/shared/repo"
	"github.com/GagulProject/go-whisky/internal/shared/scroller"
)

type WhiskyRepository interface {
	Create(context.Context, *whisky.Whisky) (*whisky.Whisky, error)
	ScrollAll(context.Context, ...scroller.ScrollOptionFn[time.Time]) (*scroller.Page[*whisky.Whisky], error)
}

type whiskyRepository struct {
	repo.Repo[*whisky.Whisky, *models.Whisky]
}

func NewWhiskyRepository(db *sql.DB) WhiskyRepository {
	return &whiskyRepository{
		repo.New[*whisky.Whisky, *models.Whisky, models.WhiskySlice](
			db,
			toBoilerFunc,
			toModelFunc,
			func(bs models.WhiskySlice) []*models.Whisky { return bs },
			queryStarter,
		),
	}
}

func (r *whiskyRepository) ScrollAll(ctx context.Context, optionFns ...scroller.ScrollOptionFn[time.Time]) (*scroller.Page[*whisky.Whisky], error) {
	option := scroller.NewScrollOption(optionFns...)

	page, err := scroller.Scroll[*whisky.Whisky, *models.Whisky](
		ctx,
		r,
		option,
		models.WhiskyColumns.CreatedAt,
	)
	return page, errors.Wrap(err)
}

func toBoilerFunc(whisky *whisky.Whisky) (*models.Whisky, error) {
	return &models.Whisky{
		Strength:  whisky.Strength,
		Size:      whisky.Size,
		CreatedAt: whisky.CreatedAt,
		UpdatedAt: whisky.CreatedAt,
	}, nil
}

func toModelFunc(boiler *models.Whisky) (*whisky.Whisky, error) {
	return &whisky.Whisky{
		Strength:  boiler.Strength,
		Size:      boiler.Size,
		CreatedAt: boiler.CreatedAt,
		UpdatedAt: boiler.UpdatedAt,
	}, nil
}

func queryStarter(
	mods ...qm.QueryMod,
) repo.Query[*models.Whisky, models.WhiskySlice] {
	return models.Whiskies(mods...)
}
