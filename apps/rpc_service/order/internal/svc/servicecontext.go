package svc

import (
	"snkrs/apps/rpc_service/order/internal/config"
	"snkrs/apps/rpc_service/order/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrderlistModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderlistModel(conn, c.CacheRedis),
	}
}
