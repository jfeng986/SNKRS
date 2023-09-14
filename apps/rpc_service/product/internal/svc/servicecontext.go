package svc

import (
	"snkrs/apps/rpc_service/product/internal/config"
	"snkrs/apps/rpc_service/product/internal/model"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/sync/singleflight"
)

type ServiceContext struct {
	Config         config.Config
	ProductModel   model.ProductModel
	CategoryModel  model.CategoryModel
	BizRedis       *redis.Redis
	SingleGroup    singleflight.Group
	LocalCache     *collection.Cache
	OperationModel model.ProductOperationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:        c,
		ProductModel:  model.NewProductModel(conn, c.CacheRedis),
		CategoryModel: model.NewCategoryModel(conn, c.CacheRedis),
		BizRedis:      redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
	}
}
