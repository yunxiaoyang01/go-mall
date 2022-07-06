package service

import (
	"github.com/yunxiaoyang01/go-mall/mall-admin/service/example"
	"github.com/yunxiaoyang01/go-mall/mall-admin/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
