import service from '@/utils/request'
// @Tags Cookie
// @Summary 创建Cookie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cookie true "创建Cookie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /JMCookie/createCookie [post]
export const createCookie = (data) => {
  return service({
    url: '/JMCookie/createCookie',
    method: 'post',
    data
  })
}

// @Tags Cookie
// @Summary 删除Cookie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cookie true "删除Cookie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /JMCookie/deleteCookie [delete]
export const deleteCookie = (params) => {
  return service({
    url: '/JMCookie/deleteCookie',
    method: 'delete',
    params
  })
}

// @Tags Cookie
// @Summary 批量删除Cookie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Cookie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /JMCookie/deleteCookie [delete]
export const deleteCookieByIds = (params) => {
  return service({
    url: '/JMCookie/deleteCookieByIds',
    method: 'delete',
    params
  })
}

// @Tags Cookie
// @Summary 更新Cookie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cookie true "更新Cookie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /JMCookie/updateCookie [put]
export const updateCookie = (data) => {
  return service({
    url: '/JMCookie/updateCookie',
    method: 'put',
    data
  })
}

// @Tags Cookie
// @Summary 用id查询Cookie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Cookie true "用id查询Cookie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /JMCookie/findCookie [get]
export const findCookie = (params) => {
  return service({
    url: '/JMCookie/findCookie',
    method: 'get',
    params
  })
}

// @Tags Cookie
// @Summary 分页获取Cookie列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Cookie列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /JMCookie/getCookieList [get]
export const getCookieList = (params) => {
  return service({
    url: '/JMCookie/getCookieList',
    method: 'get',
    params
  })
}

// @Tags Cookie
// @Summary 不需要鉴权的Cookie接口
// @accept application/json
// @Produce application/json
// @Param data query CookieReq.CookieSearch true "分页获取Cookie列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /JMCookie/getCookiePublic [get]
export const getCookiePublic = () => {
  return service({
    url: '/JMCookie/getCookiePublic',
    method: 'get',
  })
}
