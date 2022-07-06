package service

import (
	"github.com/yunxiaoyang01/go-mall/server/service/example"
	"github.com/yunxiaoyang01/go-mall/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
