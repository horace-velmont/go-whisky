package scroller

type ScrollOptionFn[S any] func(o *ScrollOption[S])

type ScrollOption[S any] struct {
	Order SortOrder
	Limit int
	Since *S
}

func NewScrollOption[S any](optFns ...ScrollOptionFn[S]) *ScrollOption[S] {
	opt := newDefaultOption[S]()
	for _, optFn := range optFns {
		optFn(opt)
	}
	return opt
}

func newDefaultOption[S any]() *ScrollOption[S] {
	return &ScrollOption[S]{
		Order: OrderDesc,
		Limit: pageDefaultLimit,
	}
}

func WithLimit[S any](limit int) ScrollOptionFn[S] {
	return func(o *ScrollOption[S]) {
		o.Limit = limit
	}
}

func WithOrder[S any](order SortOrder) ScrollOptionFn[S] {
	return func(o *ScrollOption[S]) {
		o.Order = order
	}
}

func WithSince[S any](since *S) ScrollOptionFn[S] {
	return func(o *ScrollOption[S]) {
		o.Since = since
	}
}
