package svc

import (
	"snkrs/apps/api_gateway/internal/config"

	"snkrs/apps/rpc_service/order/orderservice"
	"snkrs/apps/rpc_service/product/productservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   orderservice.OrderService
	ProductRPC productservice.ProductService
	// CommentRPC commentservice.CommentService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   orderservice.NewOrderService(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: productservice.NewProductService(zrpc.MustNewClient(c.ProductRPC)),
		// CommentRPC: commentservice.NewCommentService(zrpc.MustNewClient(c.CommentRPC)),
	}
}
