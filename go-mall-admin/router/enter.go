package router

import (
	"github.com/yunxiaoyang01/go-mall/go-mall-admin/router/example"
	"github.com/yunxiaoyang01/go-mall/go-mall-admin/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
