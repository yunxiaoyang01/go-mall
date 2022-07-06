package request

import (
	"github.com/yunxiaoyang01/go-mall/go-mall-admin/model/common/request"
	"github.com/yunxiaoyang01/go-mall/go-mall-admin/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
