package scroller

import "strings"

type SortOrder string

const (
	SortOrderAsc SortOrder = "ASC"
	OrderDesc    SortOrder = "DESC"
)

func (o SortOrder) EQ(other SortOrder) bool {
	return strings.ToUpper(string(o)) == strings.ToUpper(string(other))
}
