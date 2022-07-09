package service

import (
	"github.com/yunxiaoyang01/go-mall/mall-admin/service/example"
	"github.com/yunxiaoyang01/go-mall/mall-admin/service/product"
	"github.com/yunxiaoyang01/go-mall/mall-admin/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	ProductServiceGroup product.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
