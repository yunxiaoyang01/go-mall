package response

import (
	"github.com/yunxiaoyang01/go-mall/go-mall-admin/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
