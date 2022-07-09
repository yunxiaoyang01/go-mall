package request

import (
	"github.com/yunxiaoyang01/go-mall/mall-admin/model/product"
	"github.com/yunxiaoyang01/go-mall/mall-admin/model/common/request"
)

type PmsProductSearch struct{
    product.PmsProduct
    request.PageInfo
}
