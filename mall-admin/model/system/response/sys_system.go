package response

import "github.com/yunxiaoyang01/go-mall/mall-admin/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
