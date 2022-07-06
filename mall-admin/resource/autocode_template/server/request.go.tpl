package request

import (
	"github.com/yunxiaoyang01/go-mall/mall-admin/model/{{.Package}}"
	"github.com/yunxiaoyang01/go-mall/mall-admin/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
