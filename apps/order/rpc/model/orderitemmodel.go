package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderitemModel = (*customOrderitemModel)(nil)

type (
	// OrderitemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderitemModel.
	OrderitemModel interface {
		orderitemModel
		FindByUidOid(context.Context, int64, int64) ([]*Orderitem, error)
	}

	customOrderitemModel struct {
		*defaultOrderitemModel
	}
)

func (m *customOrderitemModel) FindByUidOid(ctx context.Context, uid int64, oid int64) ([]*Orderitem, error) {
	items := make([]*Orderitem, 0)
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `order_id` = ? order by create_time desc", orderitemRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &items, query, uid, oid)

	switch err {
	case nil:
		return items, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewOrderitemModel returns a model for the database table.
func NewOrderitemModel(conn sqlx.SqlConn, c cache.CacheConf) OrderitemModel {
	return &customOrderitemModel{
		defaultOrderitemModel: newOrderitemModel(conn, c),
	}
}
