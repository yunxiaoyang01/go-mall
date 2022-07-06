package request

import (
	"github.com/yunxiaoyang01/go-mall/server/model/{{.Package}}"
	"github.com/yunxiaoyang01/go-mall/server/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
