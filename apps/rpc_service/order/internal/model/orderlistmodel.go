package model

import (
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderlistModel = (*customOrderlistModel)(nil)

type (
	// OrderlistModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderlistModel.
	OrderlistModel interface {
		orderlistModel
		CreateOrder(ctx context.Context, oid string, uid, pid int64) error
	}

	customOrderlistModel struct {
		*defaultOrderlistModel
	}
)

// NewOrderlistModel returns a model for the database table.
func NewOrderlistModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OrderlistModel {
	return &customOrderlistModel{
		defaultOrderlistModel: newOrderlistModel(conn, c, opts...),
	}
}

func (m *customOrderlistModel) CreateOrder(ctx context.Context, oid string, uid, pid int64) error {
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		err := conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
			_, err := session.ExecCtx(ctx, "INSERT INTO orderlist(id, userid) VALUES(?,?)", oid, uid)
			if err != nil {
				return err
			}
			_, err = session.ExecCtx(ctx, "INSERT INTO orderitem(orderid, userid, proid) VALUES(?,?,?)", "", uid, pid)
			return err
		})
		return nil, err
	})
	return err
}
