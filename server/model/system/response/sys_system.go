package response

import "github.com/yunxiaoyang01/go-mall/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
