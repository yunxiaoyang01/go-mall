package system

import (
	"github.com/yunxiaoyang01/go-mall/server/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
