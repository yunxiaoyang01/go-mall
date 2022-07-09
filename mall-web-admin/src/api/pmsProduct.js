import service from '@/utils/request'

// @Tags PmsProduct
// @Summary 创建PmsProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PmsProduct true "创建PmsProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pmsProduct/createPmsProduct [post]
export const createPmsProduct = (data) => {
  return service({
    url: '/pmsProduct/createPmsProduct',
    method: 'post',
    data
  })
}

// @Tags PmsProduct
// @Summary 删除PmsProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PmsProduct true "删除PmsProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pmsProduct/deletePmsProduct [delete]
export const deletePmsProduct = (data) => {
  return service({
    url: '/pmsProduct/deletePmsProduct',
    method: 'delete',
    data
  })
}

// @Tags PmsProduct
// @Summary 删除PmsProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PmsProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pmsProduct/deletePmsProduct [delete]
export const deletePmsProductByIds = (data) => {
  return service({
    url: '/pmsProduct/deletePmsProductByIds',
    method: 'delete',
    data
  })
}

// @Tags PmsProduct
// @Summary 更新PmsProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PmsProduct true "更新PmsProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pmsProduct/updatePmsProduct [put]
export const updatePmsProduct = (data) => {
  return service({
    url: '/pmsProduct/updatePmsProduct',
    method: 'put',
    data
  })
}

// @Tags PmsProduct
// @Summary 用id查询PmsProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PmsProduct true "用id查询PmsProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pmsProduct/findPmsProduct [get]
export const findPmsProduct = (params) => {
  return service({
    url: '/pmsProduct/findPmsProduct',
    method: 'get',
    params
  })
}

// @Tags PmsProduct
// @Summary 分页获取PmsProduct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PmsProduct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pmsProduct/getPmsProductList [get]
export const getPmsProductList = (params) => {
  return service({
    url: '/pmsProduct/getPmsProductList',
    method: 'get',
    params
  })
}


