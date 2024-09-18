package Title

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Title"
    TitleReq "github.com/flipped-aurora/gin-vue-admin/server/model/Title/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type JMheathyTitlesApi struct {}



// CreateJMheathyTitles 创建标题表
// @Tags JMheathyTitles
// @Summary 创建标题表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Title.JMheathyTitles true "创建标题表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /JMTitle/createJMheathyTitles [post]
func (JMTitleApi *JMheathyTitlesApi) CreateJMheathyTitles(c *gin.Context) {
	var JMTitle Title.JMheathyTitles
	err := c.ShouldBindJSON(&JMTitle)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    JMTitle.CreatedBy = utils.GetUserID(c)
	err = JMTitleService.CreateJMheathyTitles(&JMTitle)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteJMheathyTitles 删除标题表
// @Tags JMheathyTitles
// @Summary 删除标题表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Title.JMheathyTitles true "删除标题表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /JMTitle/deleteJMheathyTitles [delete]
func (JMTitleApi *JMheathyTitlesApi) DeleteJMheathyTitles(c *gin.Context) {
	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := JMTitleService.DeleteJMheathyTitles(ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteJMheathyTitlesByIds 批量删除标题表
// @Tags JMheathyTitles
// @Summary 批量删除标题表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /JMTitle/deleteJMheathyTitlesByIds [delete]
func (JMTitleApi *JMheathyTitlesApi) DeleteJMheathyTitlesByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := JMTitleService.DeleteJMheathyTitlesByIds(IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateJMheathyTitles 更新标题表
// @Tags JMheathyTitles
// @Summary 更新标题表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Title.JMheathyTitles true "更新标题表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /JMTitle/updateJMheathyTitles [put]
func (JMTitleApi *JMheathyTitlesApi) UpdateJMheathyTitles(c *gin.Context) {
	var JMTitle Title.JMheathyTitles
	err := c.ShouldBindJSON(&JMTitle)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    JMTitle.UpdatedBy = utils.GetUserID(c)
	err = JMTitleService.UpdateJMheathyTitles(JMTitle)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindJMheathyTitles 用id查询标题表
// @Tags JMheathyTitles
// @Summary 用id查询标题表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Title.JMheathyTitles true "用id查询标题表"
// @Success 200 {object} response.Response{data=Title.JMheathyTitles,msg=string} "查询成功"
// @Router /JMTitle/findJMheathyTitles [get]
func (JMTitleApi *JMheathyTitlesApi) FindJMheathyTitles(c *gin.Context) {
	ID := c.Query("ID")
	reJMTitle, err := JMTitleService.GetJMheathyTitles(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reJMTitle, c)
}

// GetJMheathyTitlesList 分页获取标题表列表
// @Tags JMheathyTitles
// @Summary 分页获取标题表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query TitleReq.JMheathyTitlesSearch true "分页获取标题表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /JMTitle/getJMheathyTitlesList [get]
func (JMTitleApi *JMheathyTitlesApi) GetJMheathyTitlesList(c *gin.Context) {
	var pageInfo TitleReq.JMheathyTitlesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := JMTitleService.GetJMheathyTitlesInfoList(pageInfo)
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

// GetJMheathyTitlesPublic 不需要鉴权的标题表接口
// @Tags JMheathyTitles
// @Summary 不需要鉴权的标题表接口
// @accept application/json
// @Produce application/json
// @Param data query TitleReq.JMheathyTitlesSearch true "分页获取标题表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /JMTitle/getJMheathyTitlesPublic [get]
func (JMTitleApi *JMheathyTitlesApi) GetJMheathyTitlesPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    JMTitleService.GetJMheathyTitlesPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的标题表接口信息",
    }, "获取成功", c)
}
