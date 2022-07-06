package initialize

import (
	_ "github.com/yunxiaoyang01/go-mall/server/source/example"
	_ "github.com/yunxiaoyang01/go-mall/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
