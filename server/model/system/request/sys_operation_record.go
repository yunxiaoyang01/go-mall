package request

import (
	"github.com/yunxiaoyang01/go-mall/server/model/common/request"
	"github.com/yunxiaoyang01/go-mall/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
