// Code generated by goctl. DO NOT EDIT!

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
	ordersFieldNames          = builder.RawFieldNames(&Orders{})
	ordersRows                = strings.Join(ordersFieldNames, ",")
	ordersRowsExpectAutoSet   = strings.Join(stringx.Remove(ordersFieldNames, "`create_time`", "`update_time`"), ",")
	ordersRowsWithPlaceHolder = strings.Join(stringx.Remove(ordersFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheOrdersOrdersIdPrefix = "cache:orders:orders:id:"
)

type (
	ordersModel interface {
		Insert(ctx context.Context, data *Orders) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*Orders, error)
		Update(ctx context.Context, data *Orders) error
		Delete(ctx context.Context, id string) error
	}

	defaultOrdersModel struct {
		sqlc.CachedConn
		table string
	}

	Orders struct {
		Id          string    `db:"id"`          // 订单id
		Userid      int64     `db:"userid"`      // 用户id
		Shoppingid  int64     `db:"shoppingid"`  // 收货信息表id
		Payment     float64   `db:"payment"`     // 实际付款金额,单位是元,保留两位小数
		Paymenttype int64     `db:"paymenttype"` // 支付类型,1-在线支付
		Postage     int64     `db:"postage"`     // 运费,单位是元
		Status      int64     `db:"status"`      // 订单状态:0-已取消-10-未付款，20-已付款，30-待发货 40-待收货，50-交易成功，60-交易关闭
		CreateTime  time.Time `db:"create_time"` // 创建时间
		UpdateTime  time.Time `db:"update_time"` // 更新时间
	}
)

func newOrdersModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultOrdersModel {
	return &defaultOrdersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`orders`",
	}
}

func (m *defaultOrdersModel) Insert(ctx context.Context, data *Orders) (sql.Result, error) {
	ordersOrdersIdKey := fmt.Sprintf("%s%v", cacheOrdersOrdersIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, ordersRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Userid, data.Shoppingid, data.Payment, data.Paymenttype, data.Postage, data.Status)
	}, ordersOrdersIdKey)
	return ret, err
}

func (m *defaultOrdersModel) FindOne(ctx context.Context, id string) (*Orders, error) {
	ordersOrdersIdKey := fmt.Sprintf("%s%v", cacheOrdersOrdersIdPrefix, id)
	var resp Orders
	err := m.QueryRowCtx(ctx, &resp, ordersOrdersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ordersRows, m.table)
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

func (m *defaultOrdersModel) Update(ctx context.Context, data *Orders) error {
	ordersOrdersIdKey := fmt.Sprintf("%s%v", cacheOrdersOrdersIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ordersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Userid, data.Shoppingid, data.Payment, data.Paymenttype, data.Postage, data.Status, data.Id)
	}, ordersOrdersIdKey)
	return err
}

func (m *defaultOrdersModel) Delete(ctx context.Context, id string) error {
	ordersOrdersIdKey := fmt.Sprintf("%s%v", cacheOrdersOrdersIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, ordersOrdersIdKey)
	return err
}

func (m *defaultOrdersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheOrdersOrdersIdPrefix, primary)
}

func (m *defaultOrdersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ordersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultOrdersModel) tableName() string {
	return m.table
}