package Project

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Project"
	ProjectReq "github.com/flipped-aurora/gin-vue-admin/server/model/Project/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemProjectApi struct{}

// CreateSystemProject 创建项目
// @Tags SystemProject
// @Summary 创建项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Project.SystemProject true "创建项目"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /Projects/createSystemProject [post]
func (ProjectsApi *SystemProjectApi) CreateSystemProject(c *gin.Context) {
	var Projects Project.SystemProject
	err := c.ShouldBindJSON(&Projects)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Projects.CreatedBy = utils.GetUserID(c)
	err = ProjectsService.CreateSystemProject(&Projects)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSystemProject 删除项目
// @Tags SystemProject
// @Summary 删除项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Project.SystemProject true "删除项目"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /Projects/deleteSystemProject [delete]
func (ProjectsApi *SystemProjectApi) DeleteSystemProject(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := ProjectsService.DeleteSystemProject(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSystemProjectByIds 批量删除项目
// @Tags SystemProject
// @Summary 批量删除项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /Projects/deleteSystemProjectByIds [delete]
func (ProjectsApi *SystemProjectApi) DeleteSystemProjectByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := ProjectsService.DeleteSystemProjectByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSystemProject 更新项目
// @Tags SystemProject
// @Summary 更新项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Project.SystemProject true "更新项目"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /Projects/updateSystemProject [put]
func (ProjectsApi *SystemProjectApi) UpdateSystemProject(c *gin.Context) {
	var Projects Project.SystemProject
	err := c.ShouldBindJSON(&Projects)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Projects.UpdatedBy = utils.GetUserID(c)
	err = ProjectsService.UpdateSystemProject(Projects)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSystemProject 用id查询项目
// @Tags SystemProject
// @Summary 用id查询项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Project.SystemProject true "用id查询项目"
// @Success 200 {object} response.Response{data=Project.SystemProject,msg=string} "查询成功"
// @Router /Projects/findSystemProject [get]
func (ProjectsApi *SystemProjectApi) FindSystemProject(c *gin.Context) {
	ID := c.Query("ID")
	reProjects, err := ProjectsService.GetSystemProject(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reProjects, c)
}

// GetSystemProjectList 分页获取项目列表
// @Tags SystemProject
// @Summary 分页获取项目列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ProjectReq.SystemProjectSearch true "分页获取项目列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /Projects/getSystemProjectList [get]
func (ProjectsApi *SystemProjectApi) GetSystemProjectList(c *gin.Context) {
	var pageInfo ProjectReq.SystemProjectSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := ProjectsService.GetSystemProjectInfoList(pageInfo)
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

// GetSystemProjectPublic 不需要鉴权的项目接口
// @Tags SystemProject
// @Summary 不需要鉴权的项目接口
// @accept application/json
// @Produce application/json
// @Param data query ProjectReq.SystemProjectSearch true "分页获取项目列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Projects/getSystemProjectPublic [get]
func (ProjectsApi *SystemProjectApi) GetSystemProjectPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	ProjectsService.GetSystemProjectPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的项目接口信息",
	}, "获取成功", c)
}

// WriteWord 写文
// @Tags SystemProject
// @Summary 写文
// @accept application/json
// @Produce application/json
// @Param data query ProjectReq.SystemProjectSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /Projects/WriteWord [POST]
func (ProjectsApi *SystemProjectApi) WriteWord(c *gin.Context) {
	// 请添加自己的业务逻辑
	ID := c.Query("ID")
	err := ProjectsService.WriteWord(ID)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("返回数据", c)
}

// PublishArticle 发布文章
// @Tags SystemProject
// @Summary 发布文章
// @accept application/json
// @Produce application/json
// @Param data query ProjectReq.SystemProjectSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /Projects/PublishArticle [POST]
func (ProjectsApi *SystemProjectApi) PublishArticle(c *gin.Context) {
	// 请添加自己的业务逻辑
	ID := c.Query("ID")
	err := ProjectsService.PublishArticle(ID)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("返回数据", c)
}

// SyncTitle 写入标题列表
// @Tags SystemProject
// @Summary 写入标题列表
// @accept application/json
// @Produce application/json
// @Param data query ProjectReq.SystemProjectSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /Projects/SyncTitle [PUT]
func (ProjectsApi *SystemProjectApi) SyncTitle(c *gin.Context) {
	// 请添加自己的业务逻辑
	var Projects Project.SystemProject
	err := c.ShouldBindJSON(&Projects)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Projects.UpdatedBy = utils.GetUserID(c)
	err = ProjectsService.SyncTitle(Projects)
	if err != nil {
		global.GVA_LOG.Error("同步标题失败!", zap.Error(err))
		response.FailWithMessage("同步标题失败", c)
		return
	}
	response.OkWithData("同步标题成功", c)
}
