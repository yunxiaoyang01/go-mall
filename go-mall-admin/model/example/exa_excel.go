package example

import "github.com/yunxiaoyang01/go-mall/go-mall-admin/model/system"

type ExcelInfo struct {
	FileName string               `json:"fileName"` // 文件名
	InfoList []system.SysBaseMenu `json:"infoList"`
}
