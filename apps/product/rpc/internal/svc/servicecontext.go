package svc

import "mylebron/apps/product/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
	// ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
