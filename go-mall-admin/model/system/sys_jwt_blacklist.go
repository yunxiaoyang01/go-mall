package system

import (
	"github.com/yunxiaoyang01/go-mall/go-mall-admin/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
