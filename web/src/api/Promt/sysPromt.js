import service from '@/utils/request'
// @Tags Promt
// @Summary 创建提示词
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Promt true "创建提示词"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pmt/createPromt [post]
export const createPromt = (data) => {
  return service({
    url: '/pmt/createPromt',
    method: 'post',
    data
  })
}

// @Tags Promt
// @Summary 删除提示词
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Promt true "删除提示词"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pmt/deletePromt [delete]
export const deletePromt = (params) => {
  return service({
    url: '/pmt/deletePromt',
    method: 'delete',
    params
  })
}

// @Tags Promt
// @Summary 批量删除提示词
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除提示词"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pmt/deletePromt [delete]
export const deletePromtByIds = (params) => {
  return service({
    url: '/pmt/deletePromtByIds',
    method: 'delete',
    params
  })
}

// @Tags Promt
// @Summary 更新提示词
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Promt true "更新提示词"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pmt/updatePromt [put]
export const updatePromt = (data) => {
  return service({
    url: '/pmt/updatePromt',
    method: 'put',
    data
  })
}

// @Tags Promt
// @Summary 用id查询提示词
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Promt true "用id查询提示词"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pmt/findPromt [get]
export const findPromt = (params) => {
  return service({
    url: '/pmt/findPromt',
    method: 'get',
    params
  })
}

// @Tags Promt
// @Summary 分页获取提示词列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取提示词列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pmt/getPromtList [get]
export const getPromtList = (params) => {
  return service({
    url: '/pmt/getPromtList',
    method: 'get',
    params
  })
}

// @Tags Promt
// @Summary 不需要鉴权的提示词接口
// @accept application/json
// @Produce application/json
// @Param data query PromtReq.PromtSearch true "分页获取提示词列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /pmt/getPromtPublic [get]
export const getPromtPublic = () => {
  return service({
    url: '/pmt/getPromtPublic',
    method: 'get',
  })
}
