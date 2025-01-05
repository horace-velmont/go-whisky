package scroller

const (
	pageDefaultLimit    = 100
	pageDefaultOrder    = OrderDesc
	pageDefaultMaxLimit = 100
)

type PageRequest[S any] struct {
	Since *S        `query:"since" form:"since" validate:"-"`
	Limit int       `query:"limit" form:"limit" validate:"-"`
	Order SortOrder `query:"order" form:"order" validate:"-"`
}

func NewPageRequest[S any]() *PageRequest[S] {
	return &PageRequest[S]{
		Limit: pageDefaultLimit,
		Order: pageDefaultOrder,
	}
}

func (req *PageRequest[S]) ToScrollOption() *ScrollOption[S] {
	return NewScrollOption(
		WithSince[S](req.Since),
		WithLimit[S](req.Limit),
		WithOrder[S](req.Order),
	)
}

type Page[T any] struct {
	Collection []T
	Next       T
}
