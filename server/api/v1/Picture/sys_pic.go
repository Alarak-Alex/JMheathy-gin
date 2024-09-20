package Picture

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Picture"
	PictureReq "github.com/flipped-aurora/gin-vue-admin/server/model/Picture/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PictureApi struct{}

// CreatePicture 创建图片
// @Tags Picture
// @Summary 创建图片
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Picture.Picture true "创建图片"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /Pic/createPicture [post]
func (PicApi *PictureApi) CreatePicture(c *gin.Context) {
	var Pic Picture.Picture
	err := c.ShouldBindJSON(&Pic)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Pic.UpLoadUser = utils.GetUserID(c)
	Pic.CreatedBy = utils.GetUserID(c)
	err = PicService.CreatePicture(&Pic)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeletePicture 删除图片
// @Tags Picture
// @Summary 删除图片
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Picture.Picture true "删除图片"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /Pic/deletePicture [delete]
func (PicApi *PictureApi) DeletePicture(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := PicService.DeletePicture(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePictureByIds 批量删除图片
// @Tags Picture
// @Summary 批量删除图片
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /Pic/deletePictureByIds [delete]
func (PicApi *PictureApi) DeletePictureByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := PicService.DeletePictureByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePicture 更新图片
// @Tags Picture
// @Summary 更新图片
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Picture.Picture true "更新图片"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /Pic/updatePicture [put]
func (PicApi *PictureApi) UpdatePicture(c *gin.Context) {
	var Pic Picture.Picture
	err := c.ShouldBindJSON(&Pic)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Pic.UpdatedBy = utils.GetUserID(c)
	err = PicService.UpdatePicture(Pic)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPicture 用id查询图片
// @Tags Picture
// @Summary 用id查询图片
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Picture.Picture true "用id查询图片"
// @Success 200 {object} response.Response{data=Picture.Picture,msg=string} "查询成功"
// @Router /Pic/findPicture [get]
func (PicApi *PictureApi) FindPicture(c *gin.Context) {
	ID := c.Query("ID")
	rePic, err := PicService.GetPicture(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rePic, c)
}

// GetPictureList 分页获取图片列表
// @Tags Picture
// @Summary 分页获取图片列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query PictureReq.PictureSearch true "分页获取图片列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /Pic/getPictureList [get]
func (PicApi *PictureApi) GetPictureList(c *gin.Context) {
	var pageInfo PictureReq.PictureSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := PicService.GetPictureInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetPicturePublic 不需要鉴权的图片接口
// @Tags Picture
// @Summary 不需要鉴权的图片接口
// @accept application/json
// @Produce application/json
// @Param data query PictureReq.PictureSearch true "分页获取图片列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Pic/getPicturePublic [get]
func (PicApi *PictureApi) GetPicturePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	PicService.GetPicturePublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的图片接口信息",
	}, "获取成功", c)
}

// CreateMorePic 方法介绍
// @Tags Picture
// @Summary 方法介绍
// @accept application/json
// @Produce application/json
// @Param data query PictureReq.PictureSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /Pic/CreateMorePic [POST]
func (PicApi *PictureApi) CreateMorePic(c *gin.Context) {
	// 请添加自己的业务逻辑
	var PicList []Picture.Picture
	err := c.ShouldBindJSON(&PicList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	for _, Pic := range PicList {
		Pic.UpLoadUser = utils.GetUserID(c)
		Pic.CreatedBy = utils.GetUserID(c)
		err = PicService.CreatePicture(&Pic)
		if err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败:"+err.Error(), c)
			return
		}
	}
	response.OkWithMessage("创建成功", c)

	// err := PicService.CreateMorePic()
	// if err != nil {
	// 	global.GVA_LOG.Error("失败!", zap.Error(err))
	// 	response.FailWithMessage("失败", c)
	// 	return
	// }
	// response.OkWithData("返回数据", c)
}
