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
	productOperationFieldNames          = builder.RawFieldNames(&ProductOperation{})
	productOperationRows                = strings.Join(productOperationFieldNames, ",")
	productOperationRowsExpectAutoSet   = strings.Join(stringx.Remove(productOperationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	productOperationRowsWithPlaceHolder = strings.Join(stringx.Remove(productOperationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheProductProductOperationIdPrefix = "cache:product:productOperation:id:"
)

type (
	productOperationModel interface {
		Insert(ctx context.Context, data *ProductOperation) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ProductOperation, error)
		Update(ctx context.Context, data *ProductOperation) error
		Delete(ctx context.Context, id int64) error
	}

	defaultProductOperationModel struct {
		sqlc.CachedConn
		table string
	}

	ProductOperation struct {
		Id         int64     `db:"id"`
		ProductId  int64     `db:"product_id"`  // product id
		Status     int64     `db:"status"`      // 0-off market, 1-on market
		CreateTime time.Time `db:"create_time"` // creation time
		UpdateTime time.Time `db:"update_time"` // update time
	}
)

func newProductOperationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultProductOperationModel {
	return &defaultProductOperationModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`product_operation`",
	}
}

func (m *defaultProductOperationModel) withSession(session sqlx.Session) *defaultProductOperationModel {
	return &defaultProductOperationModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`product_operation`",
	}
}

func (m *defaultProductOperationModel) Delete(ctx context.Context, id int64) error {
	productProductOperationIdKey := fmt.Sprintf("%s%v", cacheProductProductOperationIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, productProductOperationIdKey)
	return err
}

func (m *defaultProductOperationModel) FindOne(ctx context.Context, id int64) (*ProductOperation, error) {
	productProductOperationIdKey := fmt.Sprintf("%s%v", cacheProductProductOperationIdPrefix, id)
	var resp ProductOperation
	err := m.QueryRowCtx(ctx, &resp, productProductOperationIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productOperationRows, m.table)
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

func (m *defaultProductOperationModel) Insert(ctx context.Context, data *ProductOperation) (sql.Result, error) {
	productProductOperationIdKey := fmt.Sprintf("%s%v", cacheProductProductOperationIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, productOperationRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ProductId, data.Status)
	}, productProductOperationIdKey)
	return ret, err
}

func (m *defaultProductOperationModel) Update(ctx context.Context, data *ProductOperation) error {
	productProductOperationIdKey := fmt.Sprintf("%s%v", cacheProductProductOperationIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, productOperationRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ProductId, data.Status, data.Id)
	}, productProductOperationIdKey)
	return err
}

func (m *defaultProductOperationModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheProductProductOperationIdPrefix, primary)
}

func (m *defaultProductOperationModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productOperationRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProductOperationModel) tableName() string {
	return m.table
}
