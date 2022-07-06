package response

import (
	"github.com/yunxiaoyang01/go-mall/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
