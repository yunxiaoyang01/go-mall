package system

import (
	"github.com/yunxiaoyang01/go-mall/go-mall-admin/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
