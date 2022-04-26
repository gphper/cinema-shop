package film

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

type PageLimitWhere struct {
	Type   int
	Cate   int
	Page   int
	Limit  int
	Status int
}

type PageLimitData []Film
