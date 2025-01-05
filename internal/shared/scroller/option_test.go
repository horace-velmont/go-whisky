package scroller_test

import (
	"testing"

	"github.com/channel-io/cht-app-commerce/internal/shared/scroller"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestNewScrollOption(t *testing.T) {
	tests := []struct {
		name        string
		optFns      []scroller.ScrollOptionFn[string]
		expectedOpt *scroller.ScrollOption[string]
	}{
		{
			name:   "default option",
			optFns: []scroller.ScrollOptionFn[string]{},
			expectedOpt: &scroller.ScrollOption[string]{
				Order: scroller.OrderDesc,
				Limit: 100,
			},
		},
		{
			name: "apply default option functions",
			optFns: []scroller.ScrollOptionFn[string]{
				scroller.WithOrder[string](scroller.OrderDesc),
				scroller.WithLimit[string](100),
			},
			expectedOpt: &scroller.ScrollOption[string]{
				Order: scroller.OrderDesc,
				Limit: 100,
			},
		},
		{
			name: "change default options",
			optFns: []scroller.ScrollOptionFn[string]{
				scroller.WithOrder[string](scroller.OrderAsc),
				scroller.WithLimit[string](20),
				scroller.WithSince(lo.ToPtr("since")),
			},
			expectedOpt: &scroller.ScrollOption[string]{
				Order: scroller.OrderAsc,
				Limit: 20,
				Since: lo.ToPtr("since"),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opt := scroller.NewScrollOption(test.optFns...)
			assert.Equal(t, test.expectedOpt, opt)
		})
	}
}
