package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mylebron/apps/app/api/internal/config"
	"mylebron/apps/order/rpc/order"
	"mylebron/apps/product/rpc/product"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   order.Order
	ProductRPC product.Product
	// ReplyRPC   reply.Reply
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   order.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: product.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		// ReplyRPC:   reply.NewReply(zrpc.MustNewClient(c.ReplyRPC)),
	}
}
