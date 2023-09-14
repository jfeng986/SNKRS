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
	shippingFieldNames          = builder.RawFieldNames(&Shipping{})
	shippingRows                = strings.Join(shippingFieldNames, ",")
	shippingRowsExpectAutoSet   = strings.Join(stringx.Remove(shippingFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	shippingRowsWithPlaceHolder = strings.Join(stringx.Remove(shippingFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheOrderShippingIdPrefix = "cache:order:shipping:id:"
)

type (
	shippingModel interface {
		Insert(ctx context.Context, data *Shipping) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Shipping, error)
		Update(ctx context.Context, data *Shipping) error
		Delete(ctx context.Context, id int64) error
	}

	defaultShippingModel struct {
		sqlc.CachedConn
		table string
	}

	Shipping struct {
		Id              int64     `db:"id"`               // Receiving Information Table ID
		Orderid         string    `db:"orderid"`          // Order ID
		Userid          int64     `db:"userid"`           // User ID
		ReceiverName    string    `db:"receiver_name"`    // Recipient name
		ReceiverPhone   string    `db:"receiver_phone"`   // Recipient landline phone
		ReceiverState   string    `db:"receiver_state"`   // Province
		ReceiverCity    string    `db:"receiver_city"`    // City
		ReceiverAddress string    `db:"receiver_address"` // Detailed address
		CreateTime      time.Time `db:"create_time"`      // Creation time
		UpdateTime      time.Time `db:"update_time"`      // Update time
	}
)

func newShippingModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultShippingModel {
	return &defaultShippingModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`shipping`",
	}
}

func (m *defaultShippingModel) withSession(session sqlx.Session) *defaultShippingModel {
	return &defaultShippingModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`shipping`",
	}
}

func (m *defaultShippingModel) Delete(ctx context.Context, id int64) error {
	orderShippingIdKey := fmt.Sprintf("%s%v", cacheOrderShippingIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, orderShippingIdKey)
	return err
}

func (m *defaultShippingModel) FindOne(ctx context.Context, id int64) (*Shipping, error) {
	orderShippingIdKey := fmt.Sprintf("%s%v", cacheOrderShippingIdPrefix, id)
	var resp Shipping
	err := m.QueryRowCtx(ctx, &resp, orderShippingIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", shippingRows, m.table)
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

func (m *defaultShippingModel) Insert(ctx context.Context, data *Shipping) (sql.Result, error) {
	orderShippingIdKey := fmt.Sprintf("%s%v", cacheOrderShippingIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, shippingRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Orderid, data.Userid, data.ReceiverName, data.ReceiverPhone, data.ReceiverState, data.ReceiverCity, data.ReceiverAddress)
	}, orderShippingIdKey)
	return ret, err
}

func (m *defaultShippingModel) Update(ctx context.Context, data *Shipping) error {
	orderShippingIdKey := fmt.Sprintf("%s%v", cacheOrderShippingIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, shippingRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Orderid, data.Userid, data.ReceiverName, data.ReceiverPhone, data.ReceiverState, data.ReceiverCity, data.ReceiverAddress, data.Id)
	}, orderShippingIdKey)
	return err
}

func (m *defaultShippingModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheOrderShippingIdPrefix, primary)
}

func (m *defaultShippingModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", shippingRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultShippingModel) tableName() string {
	return m.table
}
