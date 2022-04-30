package screen

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ScreenModel = (*customScreenModel)(nil)

type (
	// ScreenModel is an interface to be customized, add more methods here,
	// and implement the added methods in customScreenModel.
	ScreenModel interface {
		screenModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		RowCusBuilder(fields string) squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Screen, error)
		FindAllS(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Info, error)
	}

	customScreenModel struct {
		*defaultScreenModel
	}

	Info struct {
		CinemaId  int    `db:"cinema_id"`
		Price     int    `db:"price"`
		StartTime string `db:"film"`
	}
)

// NewScreenModel returns a model for the database table.
func NewScreenModel(conn sqlx.SqlConn, c cache.CacheConf) ScreenModel {
	return &customScreenModel{
		defaultScreenModel: newScreenModel(conn, c),
	}
}

// Count
func (m *defaultScreenModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultScreenModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Screen, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("screen_id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Screen
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultScreenModel) FindAllS(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Info, error) {

	rowBuilder = rowBuilder.GroupBy("cinema_id")

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Info
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// export logic
func (m *defaultScreenModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

// export logic
func (m *defaultScreenModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(screenRows).From(m.table)
}

func (m *defaultScreenModel) RowCusBuilder(fields string) squirrel.SelectBuilder {
	return squirrel.Select(fields).From(m.table)
}

// export logic
func (m *defaultScreenModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

// export logic
func (m *defaultScreenModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}
