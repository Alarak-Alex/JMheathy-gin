package Promt

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Promt"
    PromtReq "github.com/flipped-aurora/gin-vue-admin/server/model/Promt/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type PromtApi struct {}



// CreatePromt 创建提示词
// @Tags Promt
// @Summary 创建提示词
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Promt.Promt true "创建提示词"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /pmt/createPromt [post]
func (pmtApi *PromtApi) CreatePromt(c *gin.Context) {
	var pmt Promt.Promt
	err := c.ShouldBindJSON(&pmt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    pmt.CreatedBy = utils.GetUserID(c)
	err = pmtService.CreatePromt(&pmt)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeletePromt 删除提示词
// @Tags Promt
// @Summary 删除提示词
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Promt.Promt true "删除提示词"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /pmt/deletePromt [delete]
func (pmtApi *PromtApi) DeletePromt(c *gin.Context) {
	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := pmtService.DeletePromt(ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePromtByIds 批量删除提示词
// @Tags Promt
// @Summary 批量删除提示词
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /pmt/deletePromtByIds [delete]
func (pmtApi *PromtApi) DeletePromtByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := pmtService.DeletePromtByIds(IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePromt 更新提示词
// @Tags Promt
// @Summary 更新提示词
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Promt.Promt true "更新提示词"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /pmt/updatePromt [put]
func (pmtApi *PromtApi) UpdatePromt(c *gin.Context) {
	var pmt Promt.Promt
	err := c.ShouldBindJSON(&pmt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    pmt.UpdatedBy = utils.GetUserID(c)
	err = pmtService.UpdatePromt(pmt)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPromt 用id查询提示词
// @Tags Promt
// @Summary 用id查询提示词
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Promt.Promt true "用id查询提示词"
// @Success 200 {object} response.Response{data=Promt.Promt,msg=string} "查询成功"
// @Router /pmt/findPromt [get]
func (pmtApi *PromtApi) FindPromt(c *gin.Context) {
	ID := c.Query("ID")
	repmt, err := pmtService.GetPromt(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(repmt, c)
}

// GetPromtList 分页获取提示词列表
// @Tags Promt
// @Summary 分页获取提示词列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query PromtReq.PromtSearch true "分页获取提示词列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /pmt/getPromtList [get]
func (pmtApi *PromtApi) GetPromtList(c *gin.Context) {
	var pageInfo PromtReq.PromtSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := pmtService.GetPromtInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetPromtPublic 不需要鉴权的提示词接口
// @Tags Promt
// @Summary 不需要鉴权的提示词接口
// @accept application/json
// @Produce application/json
// @Param data query PromtReq.PromtSearch true "分页获取提示词列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /pmt/getPromtPublic [get]
func (pmtApi *PromtApi) GetPromtPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    pmtService.GetPromtPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的提示词接口信息",
    }, "获取成功", c)
}
