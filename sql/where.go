package sql

import (
	"strings"
)

type SqlOp string

const (
	GE SqlOp = ">="
	GR SqlOp = ">"
	EQ SqlOp = "="
	LQ SqlOp = "<="
	LR SqlOp = "<"
)

type OrAnd string

const (
	OR   OrAnd = " OR "
	AND  OrAnd = " AND "
	NONE OrAnd = ""
)

type WhereItem struct {
	Filed string
	Value string
	SqlOp
	OrAnd
}

func (da *DbData) GenWhere() string {
	s := make([]string, 0)

	for _, v := range da.Where {
		s = append(s, DbMapLeft[da.DbType])
		s = append(s, v.Filed)
		s = append(s, DbMapRight[da.DbType])
		s = append(s, string(v.SqlOp))
		s = append(s, v.Value)

		s = append(s, string(v.OrAnd))
	}
	return strings.Join(s, "")
}
