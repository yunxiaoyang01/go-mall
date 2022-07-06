package router

import (
	"github.com/yunxiaoyang01/go-mall/server/router/example"
	"github.com/yunxiaoyang01/go-mall/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
