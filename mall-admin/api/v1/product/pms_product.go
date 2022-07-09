package product

import (
	"github.com/yunxiaoyang01/go-mall/mall-admin/global"
    "github.com/yunxiaoyang01/go-mall/mall-admin/model/product"
    "github.com/yunxiaoyang01/go-mall/mall-admin/model/common/request"
    productReq "github.com/yunxiaoyang01/go-mall/mall-admin/model/product/request"
    "github.com/yunxiaoyang01/go-mall/mall-admin/model/common/response"
    "github.com/yunxiaoyang01/go-mall/mall-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PmsProductApi struct {
}

var pmsProductService = service.ServiceGroupApp.ProductServiceGroup.PmsProductService


// CreatePmsProduct 创建商品信息
// @Tags PmsProduct
// @Summary 创建商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body product.PmsProduct true "创建商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pmsProduct/createPmsProduct [post]
func (pmsProductApi *PmsProductApi) CreatePmsProduct(c *gin.Context) {
	var pmsProduct product.PmsProduct
	_ = c.ShouldBindJSON(&pmsProduct)
	if err := pmsProductService.CreatePmsProduct(pmsProduct); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePmsProduct 删除商品信息
// @Tags PmsProduct
// @Summary 删除商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body product.PmsProduct true "删除商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pmsProduct/deletePmsProduct [delete]
func (pmsProductApi *PmsProductApi) DeletePmsProduct(c *gin.Context) {
	var pmsProduct product.PmsProduct
	_ = c.ShouldBindJSON(&pmsProduct)
	if err := pmsProductService.DeletePmsProduct(pmsProduct); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePmsProductByIds 批量删除商品信息
// @Tags PmsProduct
// @Summary 批量删除PmsProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /pmsProduct/deletePmsProductByIds [delete]
func (pmsProductApi *PmsProductApi) DeletePmsProductByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := pmsProductService.DeletePmsProductByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePmsProduct 更新商品信息
// @Tags PmsProduct
// @Summary 更新PmsProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body product.PmsProduct true "更新商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pmsProduct/updatePmsProduct [put]
func (pmsProductApi *PmsProductApi) UpdatePmsProduct(c *gin.Context) {
	var pmsProduct product.PmsProduct
	_ = c.ShouldBindJSON(&pmsProduct)
	if err := pmsProductService.UpdatePmsProduct(pmsProduct); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPmsProduct 用id查询商品信息
// @Tags PmsProduct
// @Summary 用id查询PmsProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query product.PmsProduct true "用id查询商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pmsProduct/findPmsProduct [get]
func (pmsProductApi *PmsProductApi) FindPmsProduct(c *gin.Context) {
	var pmsProduct product.PmsProduct
	_ = c.ShouldBindQuery(&pmsProduct)
	if repmsProduct, err := pmsProductService.GetPmsProduct(pmsProduct.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repmsProduct": repmsProduct}, c)
	}
}

// GetPmsProductList 分页获取商品信息列表
// @Tags PmsProduct
// @Summary 分页获取PmsProduct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query productReq.PmsProductSearch true "分页获取PmsProduct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pmsProduct/getPmsProductList [get]
func (pmsProductApi *PmsProductApi) GetPmsProductList(c *gin.Context) {
	var pageInfo productReq.PmsProductSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := pmsProductService.GetPmsProductInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
