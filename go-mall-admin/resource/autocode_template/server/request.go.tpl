package request

import (
	"github.com/yunxiaoyang01/go-mall/go-mall-admin/model/{{.Package}}"
	"github.com/yunxiaoyang01/go-mall/go-mall-admin/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
