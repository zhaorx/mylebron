package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mylebron/apps/order/rpc/internal/config"
	"mylebron/apps/order/rpc/model"
)

type ServiceContext struct {
	Config         config.Config
	OrderModel     model.OrdersModel
	OrderitemModel model.OrderitemModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:         c,
		OrderModel:     model.NewOrdersModel(conn, c.CacheRedis),
		OrderitemModel: model.NewOrderitemModel(conn, c.CacheRedis),
	}
}
