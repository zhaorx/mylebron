package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mylebron/apps/product/rpc/internal/config"
	"mylebron/apps/product/rpc/internal/model"
)

type ServiceContext struct {
	Config       config.Config
	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(conn, c.CacheRedis),
	}
}
