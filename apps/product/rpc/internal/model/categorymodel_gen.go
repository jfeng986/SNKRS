// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	categoryFieldNames          = builder.RawFieldNames(&Category{})
	categoryRows                = strings.Join(categoryFieldNames, ",")
	categoryRowsExpectAutoSet   = strings.Join(stringx.Remove(categoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	categoryRowsWithPlaceHolder = strings.Join(stringx.Remove(categoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	categoryModel interface {
		Insert(ctx context.Context, data *Category) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Category, error)
		Update(ctx context.Context, data *Category) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCategoryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Category struct {
		Id         int64     `db:"id"`
		Parentid   int64     `db:"parentid"`
		Name       string    `db:"name"`
		Status     int64     `db:"status"` // 1 normal,2 delete
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newCategoryModel(conn sqlx.SqlConn) *defaultCategoryModel {
	return &defaultCategoryModel{
		conn:  conn,
		table: "`category`",
	}
}

func (m *defaultCategoryModel) withSession(session sqlx.Session) *defaultCategoryModel {
	return &defaultCategoryModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`category`",
	}
}

func (m *defaultCategoryModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCategoryModel) FindOne(ctx context.Context, id int64) (*Category, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", categoryRows, m.table)
	var resp Category
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCategoryModel) Insert(ctx context.Context, data *Category) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, categoryRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Parentid, data.Name, data.Status)
	return ret, err
}

func (m *defaultCategoryModel) Update(ctx context.Context, data *Category) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, categoryRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Parentid, data.Name, data.Status, data.Id)
	return err
}

func (m *defaultCategoryModel) tableName() string {
	return m.table
}
