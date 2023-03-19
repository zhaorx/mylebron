package logic

import (
	"context"

	"github.com/spf13/cast"
	"mylebron/apps/order/rpc/internal/svc"
	"mylebron/apps/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrdersLogic {
	return &OrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrdersLogic) Orders(in *order.OrdersRequest) (*order.OrdersResponse, error) {
	// todo: add your logic here and delete this line
	or, err := l.svcCtx.OrderModel.FindOneByUid(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	oid := cast.ToInt64(or.Id)
	orderitems_model, err := l.svcCtx.OrderitemModel.FindByUidOid(l.ctx, in.UserId, oid)
	if err != nil {
		return nil, err
	}

	orders := make([]*order.OrderItem, 0)
	for _, it := range orderitems_model {
		orders = append(orders, &order.OrderItem{
			OrderId:    it.OrderId,
			Quantity:   it.Quantity,
			Payment:    0,
			ProductId:  it.ProductId,
			UserId:     it.UserId,
			CreateTime: cast.ToInt64(it.CreateTime),
		})
	}

	return &order.OrdersResponse{Orders: orders}, nil

}
