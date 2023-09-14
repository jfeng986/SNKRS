package service

import (
	"context"
	"encoding/json"

	"snkrs/apps/rpc_service/order/orderservice"
	"snkrs/apps/rpc_service/product/productservice"
	"snkrs/apps/rpc_service/reserve/rmq/internal/config"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/zrpc"
)

type Service struct {
	c          config.Config
	ProductRPC productservice.ProductService
	OrderRPC   orderservice.OrderService
}

type KafkaData struct {
	Uid int64 `json:"uid"`
	Pid int64 `json:"pid"`
}

func NewService(c config.Config) *Service {
	return &Service{
		c:          c,
		ProductRPC: productservice.NewProductService(zrpc.MustNewClient(c.ProductRPC)),
		OrderRPC:   orderservice.NewOrderService(zrpc.MustNewClient(c.OrderRPC)),
	}
}

func (s *Service) Consume(_ string, value string) error {
	logx.Infof("Consume value: %s\n", value)
	var data KafkaData
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}
	p, err := s.ProductRPC.Product(context.Background(), &productservice.ProductItemRequest{ProductId: data.Pid})
	if err != nil {
		return err
	}
	if p.Stock <= 0 {
		return nil
	}
	_, err = s.OrderRPC.CreateOrder(context.Background(), &orderservice.CreateOrderRequest{Uid: data.Uid, Pid: data.Pid})
	if err != nil {
		logx.Errorf("CreateOrder uid: %d pid: %d error: %v", data.Uid, data.Pid, err)
		return err
	}
	_, err = s.ProductRPC.UpdateProductStock(context.Background(), &productservice.UpdateProductStockRequest{ProductId: data.Pid, Num: 1})
	if err != nil {
		logx.Errorf("UpdateProductStock uid: %d pid: %d error: %v", data.Uid, data.Pid, err)
		return err
	}
	return nil
}
