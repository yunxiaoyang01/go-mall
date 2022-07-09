package router

import (
	"github.com/yunxiaoyang01/go-mall/mall-admin/router/example"
	"github.com/yunxiaoyang01/go-mall/mall-admin/router/product"
	"github.com/yunxiaoyang01/go-mall/mall-admin/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Product product.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
