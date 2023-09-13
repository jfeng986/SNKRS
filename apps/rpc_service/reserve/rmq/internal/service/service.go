package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"snkrs/apps/rpc_service/order/order"
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
	waiter     sync.WaitGroup
	msgsChan   []chan *KafkaData
}

type KafkaData struct {
	Uid int64 `json:"uid"`
	Pid int64 `json:"pid"`
}

const (
	chanCount   = 10
	bufferCount = 1024
)

func NewService(c config.Config) *Service {
	s := &Service{
		c:          c,
		ProductRPC: productservice.NewProductService(zrpc.MustNewClient(c.ProductRPC)),
		OrderRPC:   orderservice.NewOrderService(zrpc.MustNewClient(c.OrderRPC)),
		msgsChan:   make([]chan *KafkaData, chanCount),
	}
	for i := 0; i < chanCount; i++ {
		ch := make(chan *KafkaData, bufferCount)
		s.msgsChan[i] = ch
		s.waiter.Add(1)
		go s.consume(ch)
	}

	return s
}

func (s *Service) consume(ch chan *KafkaData) {
	defer s.waiter.Done()

	for {
		m, ok := <-ch
		if !ok {
			log.Fatal("reserve rmq exit")
		}
		fmt.Printf("consume msg: %+v\n", m)
		_, err := s.ProductRPC.CheckAndUpdateStock(context.Background(), &productservice.CheckAndUpdateStockRequest{ProductId: m.Pid})
		if err != nil {
			logx.Errorf("s.ProductRPC.CheckAndUpdateStock pid: %d error: %v", m.Pid, err)
			return
		}
		_, err = s.OrderRPC.CreateOrder(context.Background(), &order.CreateOrderRequest{Uid: m.Uid, Pid: m.Pid})
		if err != nil {
			logx.Errorf("CreateOrder uid: %d pid: %d error: %v", m.Uid, m.Pid, err)
			return
		}
		_, err = s.ProductRPC.UpdateProductStock(context.Background(), &productservice.UpdateProductStockRequest{ProductId: m.Pid, Num: 1})
		if err != nil {
			logx.Errorf("UpdateProductStock uid: %d pid: %d error: %v", m.Uid, m.Pid, err)
		}
	}
}

func (s *Service) Consume(_ string, value string) error {
	logx.Infof("Consume value: %s\n", value)
	var data []*KafkaData
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}
	for _, d := range data {
		s.msgsChan[d.Pid%chanCount] <- d
	}
	return nil
}
