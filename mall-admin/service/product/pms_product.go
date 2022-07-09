package product

import (
	"github.com/yunxiaoyang01/go-mall/mall-admin/global"
	"github.com/yunxiaoyang01/go-mall/mall-admin/model/product"
	"github.com/yunxiaoyang01/go-mall/mall-admin/model/common/request"
    productReq "github.com/yunxiaoyang01/go-mall/mall-admin/model/product/request"
)

type PmsProductService struct {
}

// CreatePmsProduct 创建商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (pmsProductService *PmsProductService) CreatePmsProduct(pmsProduct product.PmsProduct) (err error) {
	err = global.GVA_DB.Create(&pmsProduct).Error
	return err
}

// DeletePmsProduct 删除商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (pmsProductService *PmsProductService)DeletePmsProduct(pmsProduct product.PmsProduct) (err error) {
	err = global.GVA_DB.Delete(&pmsProduct).Error
	return err
}

// DeletePmsProductByIds 批量删除商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (pmsProductService *PmsProductService)DeletePmsProductByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]product.PmsProduct{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePmsProduct 更新商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (pmsProductService *PmsProductService)UpdatePmsProduct(pmsProduct product.PmsProduct) (err error) {
	err = global.GVA_DB.Save(&pmsProduct).Error
	return err
}

// GetPmsProduct 根据id获取商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (pmsProductService *PmsProductService)GetPmsProduct(id uint) (pmsProduct product.PmsProduct, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&pmsProduct).Error
	return
}

// GetPmsProductInfoList 分页获取商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (pmsProductService *PmsProductService)GetPmsProductInfoList(info productReq.PmsProductSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&product.PmsProduct{})
    var pmsProducts []product.PmsProduct
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&pmsProducts).Error
	return  pmsProducts, total, err
}
