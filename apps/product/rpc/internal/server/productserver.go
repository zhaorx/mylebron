// Code generated by goctl. DO NOT EDIT!
// Source: product.proto

package server

import (
	"context"

	"mylebron/apps/product/rpc/internal/logic"
	"mylebron/apps/product/rpc/internal/svc"
	"mylebron/apps/product/rpc/product"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	product.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductServer) Products(ctx context.Context, in *product.ProductRequest) (*product.ProductResponse, error) {
	l := logic.NewProductsLogic(ctx, s.svcCtx)
	return l.Products(in)
}