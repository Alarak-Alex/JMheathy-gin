import service from '@/utils/request'
// @Tags Picture
// @Summary 创建图片
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Picture true "创建图片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /Pic/createPicture [post]
export const createPicture = (data) => {
  return service({
    url: '/Pic/createPicture',
    method: 'post',
    data
  })
}

// @Tags Picture
// @Summary 删除图片
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Picture true "删除图片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Pic/deletePicture [delete]
export const deletePicture = (params) => {
  return service({
    url: '/Pic/deletePicture',
    method: 'delete',
    params
  })
}

// @Tags Picture
// @Summary 批量删除图片
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除图片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Pic/deletePicture [delete]
export const deletePictureByIds = (params) => {
  return service({
    url: '/Pic/deletePictureByIds',
    method: 'delete',
    params
  })
}

// @Tags Picture
// @Summary 更新图片
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Picture true "更新图片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Pic/updatePicture [put]
export const updatePicture = (data) => {
  return service({
    url: '/Pic/updatePicture',
    method: 'put',
    data
  })
}

// @Tags Picture
// @Summary 用id查询图片
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Picture true "用id查询图片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Pic/findPicture [get]
export const findPicture = (params) => {
  return service({
    url: '/Pic/findPicture',
    method: 'get',
    params
  })
}

// @Tags Picture
// @Summary 分页获取图片列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取图片列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Pic/getPictureList [get]
export const getPictureList = (params) => {
  return service({
    url: '/Pic/getPictureList',
    method: 'get',
    params
  })
}

// @Tags Picture
// @Summary 不需要鉴权的图片接口
// @accept application/json
// @Produce application/json
// @Param data query PictureReq.PictureSearch true "分页获取图片列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Pic/getPicturePublic [get]
export const getPicturePublic = () => {
  return service({
    url: '/Pic/getPicturePublic',
    method: 'get',
  })
}
// CreateMorePic 方法介绍
// @Tags Picture
// @Summary 方法介绍
// @accept application/json
// @Produce application/json
// @Param data query PictureReq.PictureSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /Pic/CreateMorePic [POST]
export const CreateMorePic = () => {
  return service({
    url: '/Pic/CreateMorePic',
    method: 'POST'
  })
}
