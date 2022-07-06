package v1

import (
	"github.com/yunxiaoyang01/go-mall/server/api/v1/example"
	"github.com/yunxiaoyang01/go-mall/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
