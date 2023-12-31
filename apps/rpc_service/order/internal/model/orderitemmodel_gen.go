// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	orderitemFieldNames          = builder.RawFieldNames(&Orderitem{})
	orderitemRows                = strings.Join(orderitemFieldNames, ",")
	orderitemRowsExpectAutoSet   = strings.Join(stringx.Remove(orderitemFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	orderitemRowsWithPlaceHolder = strings.Join(stringx.Remove(orderitemFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheOrderOrderitemIdPrefix = "cache:order:orderitem:id:"
)

type (
	orderitemModel interface {
		Insert(ctx context.Context, data *Orderitem) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Orderitem, error)
		Update(ctx context.Context, data *Orderitem) error
		Delete(ctx context.Context, id int64) error
	}

	defaultOrderitemModel struct {
		sqlc.CachedConn
		table string
	}

	Orderitem struct {
		Id         int64     `db:"id"`          // Order Subtable ID
		Orderid    string    `db:"orderid"`     // Order ID
		Userid     int64     `db:"userid"`      // User ID
		Proid      int64     `db:"proid"`       // Product ID
		Proname    string    `db:"proname"`     // Product name
		Proimage   string    `db:"proimage"`    // Product image address
		Price      float64   `db:"price"`       // Product price
		Quantity   int64     `db:"quantity"`    // Product quantity
		Totalprice float64   `db:"totalprice"`  // Total product price
		CreateTime time.Time `db:"create_time"` // Creation time
		UpdateTime time.Time `db:"update_time"` // Update time
	}
)

func newOrderitemModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultOrderitemModel {
	return &defaultOrderitemModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`orderitem`",
	}
}

func (m *defaultOrderitemModel) withSession(session sqlx.Session) *defaultOrderitemModel {
	return &defaultOrderitemModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`orderitem`",
	}
}

func (m *defaultOrderitemModel) Delete(ctx context.Context, id int64) error {
	orderOrderitemIdKey := fmt.Sprintf("%s%v", cacheOrderOrderitemIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, orderOrderitemIdKey)
	return err
}

func (m *defaultOrderitemModel) FindOne(ctx context.Context, id int64) (*Orderitem, error) {
	orderOrderitemIdKey := fmt.Sprintf("%s%v", cacheOrderOrderitemIdPrefix, id)
	var resp Orderitem
	err := m.QueryRowCtx(ctx, &resp, orderOrderitemIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderitemRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderitemModel) Insert(ctx context.Context, data *Orderitem) (sql.Result, error) {
	orderOrderitemIdKey := fmt.Sprintf("%s%v", cacheOrderOrderitemIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, orderitemRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Orderid, data.Userid, data.Proid, data.Proname, data.Proimage, data.Price, data.Quantity, data.Totalprice)
	}, orderOrderitemIdKey)
	return ret, err
}

func (m *defaultOrderitemModel) Update(ctx context.Context, data *Orderitem) error {
	orderOrderitemIdKey := fmt.Sprintf("%s%v", cacheOrderOrderitemIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderitemRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Orderid, data.Userid, data.Proid, data.Proname, data.Proimage, data.Price, data.Quantity, data.Totalprice, data.Id)
	}, orderOrderitemIdKey)
	return err
}

func (m *defaultOrderitemModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheOrderOrderitemIdPrefix, primary)
}

func (m *defaultOrderitemModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderitemRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultOrderitemModel) tableName() string {
	return m.table
}
