import service from '@/utils/request'
// @Tags SystemProject
// @Summary 创建项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SystemProject true "创建项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /Projects/createSystemProject [post]
export const createSystemProject = (data) => {
  return service({
    url: '/Projects/createSystemProject',
    method: 'post',
    data
  })
}

// @Tags SystemProject
// @Summary 删除项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SystemProject true "删除项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Projects/deleteSystemProject [delete]
export const deleteSystemProject = (params) => {
  return service({
    url: '/Projects/deleteSystemProject',
    method: 'delete',
    params
  })
}

// @Tags SystemProject
// @Summary 批量删除项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Projects/deleteSystemProject [delete]
export const deleteSystemProjectByIds = (params) => {
  return service({
    url: '/Projects/deleteSystemProjectByIds',
    method: 'delete',
    params
  })
}

// @Tags SystemProject
// @Summary 更新项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SystemProject true "更新项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Projects/updateSystemProject [put]
export const updateSystemProject = (data) => {
  return service({
    url: '/Projects/updateSystemProject',
    method: 'put',
    data
  })
}

// @Tags SystemProject
// @Summary 用id查询项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SystemProject true "用id查询项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Projects/findSystemProject [get]
export const findSystemProject = (params) => {
  return service({
    url: '/Projects/findSystemProject',
    method: 'get',
    params
  })
}

// @Tags SystemProject
// @Summary 分页获取项目列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取项目列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Projects/getSystemProjectList [get]
export const getSystemProjectList = (params) => {
  return service({
    url: '/Projects/getSystemProjectList',
    method: 'get',
    params
  })
}

// @Tags SystemProject
// @Summary 不需要鉴权的项目接口
// @accept application/json
// @Produce application/json
// @Param data query ProjectReq.SystemProjectSearch true "分页获取项目列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Projects/getSystemProjectPublic [get]
export const getSystemProjectPublic = () => {
  return service({
    url: '/Projects/getSystemProjectPublic',
    method: 'get',
  })
}
// WriteWord 写文
// @Tags SystemProject
// @Summary 写文
// @accept application/json
// @Produce application/json
// @Param data query ProjectReq.SystemProjectSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /Projects/WriteWord [POST]
export const WriteWord = () => {
  return service({
    url: '/Projects/WriteWord',
    method: 'POST'
  })
}
// PublishArticle 发布文章
// @Tags SystemProject
// @Summary 发布文章
// @accept application/json
// @Produce application/json
// @Param data query ProjectReq.SystemProjectSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /Projects/PublishArticle [POST]
export const PublishArticle = () => {
  return service({
    url: '/Projects/PublishArticle',
    method: 'POST'
  })
}
