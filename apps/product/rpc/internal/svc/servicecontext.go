package svc

import (
	"snkrs/apps/product/rpc/internal/config"
	"snkrs/apps/product/rpc/internal/model"
)

type ServiceContext struct {
	Config        config.Config
	ProductModel  model.ProductModel
	CategoryModel model.CategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
