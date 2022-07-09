package product

import (
	"github.com/yunxiaoyang01/go-mall/mall-admin/api/v1"
	"github.com/yunxiaoyang01/go-mall/mall-admin/middleware"
	"github.com/gin-gonic/gin"
)

type PmsProductRouter struct {
}

// InitPmsProductRouter 初始化 商品信息 路由信息
func (s *PmsProductRouter) InitPmsProductRouter(Router *gin.RouterGroup) {
	pmsProductRouter := Router.Group("pmsProduct").Use(middleware.OperationRecord())
	pmsProductRouterWithoutRecord := Router.Group("pmsProduct")
	var pmsProductApi = v1.ApiGroupApp.ProductApiGroup.PmsProductApi
	{
		pmsProductRouter.POST("createPmsProduct", pmsProductApi.CreatePmsProduct)   // 新建商品信息
		pmsProductRouter.DELETE("deletePmsProduct", pmsProductApi.DeletePmsProduct) // 删除商品信息
		pmsProductRouter.DELETE("deletePmsProductByIds", pmsProductApi.DeletePmsProductByIds) // 批量删除商品信息
		pmsProductRouter.PUT("updatePmsProduct", pmsProductApi.UpdatePmsProduct)    // 更新商品信息
	}
	{
		pmsProductRouterWithoutRecord.GET("findPmsProduct", pmsProductApi.FindPmsProduct)        // 根据ID获取商品信息
		pmsProductRouterWithoutRecord.GET("getPmsProductList", pmsProductApi.GetPmsProductList)  // 获取商品信息列表
	}
}
