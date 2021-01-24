package sql

type DbData struct {
	Table     string
	Fileds    string
	Where     []WhereItem
	OrderBy   []OrderByItem
	PageIndex int
	PageSize  int
	DbType
}
