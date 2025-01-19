package scroller

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/GagulProject/go-whisky/internal/shared/errors"
	"github.com/GagulProject/go-whisky/internal/shared/repo"
)

type ScrollFetchFn[T any] func(
	context.Context,
	...qm.QueryMod,
) ([]*T, error)

func Scroll[D repo.DomainModel, B repo.BoilerModel, A any](
	ctx context.Context,
	repo repo.Repo[D, B],
	option *ScrollOption[A],
	field string,
	mods ...qm.QueryMod,
) (*Page[D], error) {
	mods = append(
		mods,
		since(field, option),
		qm.OrderBy(fmt.Sprintf("%s %s", field, option.Order)),
		qm.Limit(option.Limit+1),
	)
	mods = lo.Filter(mods, func(item qm.QueryMod, _ int) bool {
		return item != nil
	})
	collection, err := repo.FindAllBy(
		ctx,
		mods...,
	)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	var next D
	if len(collection) > option.Limit {
		next = collection[option.Limit]
		collection = collection[:option.Limit]
	}

	return &Page[D]{
		Collection: collection,
		Next:       next,
	}, nil
}

func since[S any](field string, option *ScrollOption[S]) qm.QueryMod {
	if option.Since == nil {
		return nil
	}

	if SortOrderAsc.EQ(option.Order) {
		return qm.Where(fmt.Sprintf("%s >= ?", field), option.Since)
	} else if OrderDesc.EQ(option.Order) {
		return qm.Where(fmt.Sprintf("%s <= ?", field), option.Since)
	}
	return nil
}
