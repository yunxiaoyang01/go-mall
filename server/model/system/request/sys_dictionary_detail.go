package request

import (
	"github.com/yunxiaoyang01/go-mall/server/model/common/request"
	"github.com/yunxiaoyang01/go-mall/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
