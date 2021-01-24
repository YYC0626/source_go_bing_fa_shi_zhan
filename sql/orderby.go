package sql

import (
	"strings"
)

type Order string

const (
	ASC  Order = " ASC "
	DESC Order = " DESC "
)

type OrderByItem struct {
	Filed string
	Order
}

func (da *DbData) GenOrderBy() string {
	s := make([]string, 0)

	for _, v := range da.OrderBy {
		s = append(s, DbMapLeft[da.DbType])
		s = append(s, v.Filed)
		s = append(s, DbMapRight[da.DbType])
		s = append(s, string(v.Order))
		s = append(s, ",")
	}
	s = s[0 : len(s)-1]
	return strings.Join(s, "")
}
