package Cookie

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Cookie"
    CookieReq "github.com/flipped-aurora/gin-vue-admin/server/model/Cookie/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type CookieApi struct {}



// CreateCookie 创建Cookie
// @Tags Cookie
// @Summary 创建Cookie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Cookie.Cookie true "创建Cookie"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /JMCookie/createCookie [post]
func (JMCookieApi *CookieApi) CreateCookie(c *gin.Context) {
	var JMCookie Cookie.Cookie
	err := c.ShouldBindJSON(&JMCookie)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    JMCookie.CreatedBy = utils.GetUserID(c)
	err = JMCookieService.CreateCookie(&JMCookie)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteCookie 删除Cookie
// @Tags Cookie
// @Summary 删除Cookie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Cookie.Cookie true "删除Cookie"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /JMCookie/deleteCookie [delete]
func (JMCookieApi *CookieApi) DeleteCookie(c *gin.Context) {
	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := JMCookieService.DeleteCookie(ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCookieByIds 批量删除Cookie
// @Tags Cookie
// @Summary 批量删除Cookie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /JMCookie/deleteCookieByIds [delete]
func (JMCookieApi *CookieApi) DeleteCookieByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := JMCookieService.DeleteCookieByIds(IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCookie 更新Cookie
// @Tags Cookie
// @Summary 更新Cookie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Cookie.Cookie true "更新Cookie"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /JMCookie/updateCookie [put]
func (JMCookieApi *CookieApi) UpdateCookie(c *gin.Context) {
	var JMCookie Cookie.Cookie
	err := c.ShouldBindJSON(&JMCookie)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    JMCookie.UpdatedBy = utils.GetUserID(c)
	err = JMCookieService.UpdateCookie(JMCookie)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCookie 用id查询Cookie
// @Tags Cookie
// @Summary 用id查询Cookie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Cookie.Cookie true "用id查询Cookie"
// @Success 200 {object} response.Response{data=Cookie.Cookie,msg=string} "查询成功"
// @Router /JMCookie/findCookie [get]
func (JMCookieApi *CookieApi) FindCookie(c *gin.Context) {
	ID := c.Query("ID")
	reJMCookie, err := JMCookieService.GetCookie(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reJMCookie, c)
}

// GetCookieList 分页获取Cookie列表
// @Tags Cookie
// @Summary 分页获取Cookie列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query CookieReq.CookieSearch true "分页获取Cookie列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /JMCookie/getCookieList [get]
func (JMCookieApi *CookieApi) GetCookieList(c *gin.Context) {
	var pageInfo CookieReq.CookieSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := JMCookieService.GetCookieInfoList(pageInfo)
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

// GetCookiePublic 不需要鉴权的Cookie接口
// @Tags Cookie
// @Summary 不需要鉴权的Cookie接口
// @accept application/json
// @Produce application/json
// @Param data query CookieReq.CookieSearch true "分页获取Cookie列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /JMCookie/getCookiePublic [get]
func (JMCookieApi *CookieApi) GetCookiePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    JMCookieService.GetCookiePublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的Cookie接口信息",
    }, "获取成功", c)
}
