// 缺失mysql.go文件，无法通过编译

package main

import (
	"fmt"

	"source_go_bing_fa_shi_zhan/sql"
)

func main() {
	db := sql.DbData{}
	da.Table = "Student"
	da.Fields = "ID,NAME,AGE"

	where := make([]sql.WhereItem, 0)
	where = append(where, sql.WhereItem{
		Filed: "XH",
		Value: "'0906'",
		SqlOp: sql.EQ,
		OrAnd: sql.AND,
	})
	db.Where = where
	orderby := make([]sql.OrderByItem, 0)
	orderby = append(orderby, sql.OrderByItem{
		Filed: "XH",
		Order: sql.ASC,
	})
	db.Where = where
	db.OrderBy = orderby
	db.DbType = sql.MYSQL

	sqlBuilder := sql.SqlBuilder{}
	sql0 := sqlNuilder.GenSelectSQL(&db)
	fmt.Println(sql0)

	db.PageIndex = 1
	da.PageSize = 20
	sql0 = sqlBuilder.GenPageSQL(&db)
	fmt.Println(sql0)

	sql0 = sqlBuilder.GenSelectSQL(&db)
	fmt.Println(sql0)

	sql0 = sqlBuilder.GenPageSQL(&db)
	fmt.Println(sql0)
}
