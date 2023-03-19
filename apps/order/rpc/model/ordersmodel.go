package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrdersModel = (*customOrdersModel)(nil)

type (
	// OrdersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrdersModel.
	OrdersModel interface {
		ordersModel
		FindOneByUid(ctx context.Context, id int64) (*Orders, error)
	}

	customOrdersModel struct {
		*defaultOrdersModel
	}
)

func (m *customOrdersModel) FindOneByUid(ctx context.Context, uid int64) (*Orders, error) {
	var resp Orders
	query := fmt.Sprintf("select %s from %s where `userid` = ? order by create_time desc limit 1", ordersRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, uid)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewOrdersModel returns a model for the database table.
func NewOrdersModel(conn sqlx.SqlConn, c cache.CacheConf) OrdersModel {
	return &customOrdersModel{
		defaultOrdersModel: newOrdersModel(conn, c),
	}
}
