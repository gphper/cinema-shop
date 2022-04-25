package film

import (
	"context"
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FilmModel = (*customFilmModel)(nil)

type (
	// FilmModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFilmModel.
	FilmModel interface {
		filmModel
		PageLimit(ctx context.Context, where PageLimitWhere, count *int, data *PageLimitData) error
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

// 分页获取数据
func (m *customFilmModel) PageLimit(ctx context.Context, where PageLimitWhere, count *int, data *PageLimitData) error {

	sql := "select * from " + m.table + " where film_id > 0"
	sqlnum := "select count(`film_id`) from " + m.table + " where film_id > 0"
	if where.Cate > 0 {
		sql += " and cate =" + strconv.Itoa(where.Cate)
		sqlnum += " and cate =" + strconv.Itoa(where.Cate)
	}

	if where.Type > 0 {
		sql += " and type =" + strconv.Itoa(where.Cate)
		sqlnum += " and type =" + strconv.Itoa(where.Cate)
	}

	sql += " limit ?,?"

	key := fmt.Sprintf("film:data:%d%d%d%d", where.Cate, where.Type, where.Page, where.Limit)
	if err := m.QueryRowCtx(ctx, data, key, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		err := conn.QueryRowsCtx(ctx, data, sql, ((where.Page - 1) * where.Limit), where.Limit)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	keyCount := fmt.Sprintf("film:count:%d%d%d%d", where.Cate, where.Type, where.Page, where.Limit)
	if err := m.QueryRowCtx(ctx, count, keyCount, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		err := conn.QueryRowCtx(ctx, count, sqlnum)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
