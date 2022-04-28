package film

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FilmModel = (*customFilmModel)(nil)

type (
	// FilmModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFilmModel.
	FilmModel interface {
		filmModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, limit int64, orderBy string) ([]*Film, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
	}

	customFilmModel struct {
		*defaultFilmModel
	}
)

// NewFilmModel returns a model for the database table.
func NewFilmModel(conn sqlx.SqlConn, c cache.CacheConf) FilmModel {
	return &customFilmModel{
		defaultFilmModel: newFilmModel(conn, c),
	}
}

// 分页获取影片信息
func (m *defaultFilmModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, limit int64, orderBy string) ([]*Film, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("film_id ASC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	query, values, err := rowBuilder.Offset(uint64(offset)).Limit(uint64(limit)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Film
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultFilmModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := countBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

// export logic
func (m *defaultFilmModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

// export logic
func (m *defaultFilmModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(filmRows).From(m.table)
}

// export logic
func (m *defaultFilmModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

// export logic
func (m *defaultFilmModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}
