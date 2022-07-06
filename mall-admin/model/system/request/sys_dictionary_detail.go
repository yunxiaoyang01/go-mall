package request

import (
	"github.com/yunxiaoyang01/go-mall/mall-admin/model/common/request"
	"github.com/yunxiaoyang01/go-mall/mall-admin/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
