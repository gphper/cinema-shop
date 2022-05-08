package tickets

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TicketsModel = (*customTicketsModel)(nil)

type (
	// TicketsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTicketsModel.
	TicketsModel interface {
		ticketsModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Tickets, error)
		UpdateByOrderId(ctx context.Context, session sqlx.Session, order_id int64, data *Tickets) error
	}

	customTicketsModel struct {
		*defaultTicketsModel
	}
)

// NewTicketsModel returns a model for the database table.
func NewTicketsModel(conn sqlx.SqlConn, c cache.CacheConf) TicketsModel {
	return &customTicketsModel{
		defaultTicketsModel: newTicketsModel(conn, c),
	}
}

// Count
func (m *defaultTicketsModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultTicketsModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Tickets, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("ticket_id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Tickets
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTicketsModel) UpdateByOrderId(ctx context.Context, session sqlx.Session, order_id int64, data *Tickets) error {

	fmt.Printf("%+v", *data)

	t := reflect.TypeOf(*data)
	v := reflect.ValueOf(*data)
	sqlSlice := make([]string, 0)
	params := make([]interface{}, 0)
	for k := 0; k < t.NumField(); k++ {

		if v.Field(k).Kind() == reflect.Struct {
			if v.Field(k).FieldByName("Valid").Bool() {
				sqlSlice = append(sqlSlice, fmt.Sprintf("%s = ?", t.Field(k).Tag.Get("db")))
				params = append(params, v.Field(k).Field(0).Interface())
			}
			continue
		}

		if !v.Field(k).IsZero() {
			sqlSlice = append(sqlSlice, fmt.Sprintf("%s = ?", t.Field(k).Tag.Get("db")))
			params = append(params, v.Field(k).Interface())
		}

	}
	sqlStr := strings.Join(sqlSlice, ",")

	ticketsTicketIdKey := fmt.Sprintf("%s%v", cacheTicketsTicketIdPrefix, data.TicketId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `order_id` = %d", m.table, sqlStr, data.OrderId.Int64)
		if session != nil {
			return session.ExecCtx(ctx, query, params...)
		}
		return conn.ExecCtx(ctx, query, params...)
	}, ticketsTicketIdKey)
	return err
}

// export logic
func (m *defaultTicketsModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

// export logic
func (m *defaultTicketsModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(ticketsRows).From(m.table)
}

// export logic
func (m *defaultTicketsModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

// export logic
func (m *defaultTicketsModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}
