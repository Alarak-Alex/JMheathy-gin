import service from '@/utils/request'
// @Tags JMheathyTitles
// @Summary 创建标题表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JMheathyTitles true "创建标题表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /JMTitle/createJMheathyTitles [post]
export const createJMheathyTitles = (data) => {
  return service({
    url: '/JMTitle/createJMheathyTitles',
    method: 'post',
    data
  })
}

// @Tags JMheathyTitles
// @Summary 删除标题表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JMheathyTitles true "删除标题表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /JMTitle/deleteJMheathyTitles [delete]
export const deleteJMheathyTitles = (params) => {
  return service({
    url: '/JMTitle/deleteJMheathyTitles',
    method: 'delete',
    params
  })
}

// @Tags JMheathyTitles
// @Summary 批量删除标题表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除标题表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /JMTitle/deleteJMheathyTitles [delete]
export const deleteJMheathyTitlesByIds = (params) => {
  return service({
    url: '/JMTitle/deleteJMheathyTitlesByIds',
    method: 'delete',
    params
  })
}

// @Tags JMheathyTitles
// @Summary 更新标题表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JMheathyTitles true "更新标题表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /JMTitle/updateJMheathyTitles [put]
export const updateJMheathyTitles = (data) => {
  return service({
    url: '/JMTitle/updateJMheathyTitles',
    method: 'put',
    data
  })
}

// @Tags JMheathyTitles
// @Summary 用id查询标题表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.JMheathyTitles true "用id查询标题表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /JMTitle/findJMheathyTitles [get]
export const findJMheathyTitles = (params) => {
  return service({
    url: '/JMTitle/findJMheathyTitles',
    method: 'get',
    params
  })
}

// @Tags JMheathyTitles
// @Summary 分页获取标题表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取标题表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /JMTitle/getJMheathyTitlesList [get]
export const getJMheathyTitlesList = (params) => {
  return service({
    url: '/JMTitle/getJMheathyTitlesList',
    method: 'get',
    params
  })
}

// @Tags JMheathyTitles
// @Summary 不需要鉴权的标题表接口
// @accept application/json
// @Produce application/json
// @Param data query TitleReq.JMheathyTitlesSearch true "分页获取标题表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /JMTitle/getJMheathyTitlesPublic [get]
export const getJMheathyTitlesPublic = () => {
  return service({
    url: '/JMTitle/getJMheathyTitlesPublic',
    method: 'get',
  })
}
