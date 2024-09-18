package Title

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type JMheathyTitlesRouter struct {}

// InitJMheathyTitlesRouter 初始化 标题表 路由信息
func (s *JMheathyTitlesRouter) InitJMheathyTitlesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	JMTitleRouter := Router.Group("JMTitle").Use(middleware.OperationRecord())
	JMTitleRouterWithoutRecord := Router.Group("JMTitle")
	JMTitleRouterWithoutAuth := PublicRouter.Group("JMTitle")
	{
		JMTitleRouter.POST("createJMheathyTitles", JMTitleApi.CreateJMheathyTitles)   // 新建标题表
		JMTitleRouter.DELETE("deleteJMheathyTitles", JMTitleApi.DeleteJMheathyTitles) // 删除标题表
		JMTitleRouter.DELETE("deleteJMheathyTitlesByIds", JMTitleApi.DeleteJMheathyTitlesByIds) // 批量删除标题表
		JMTitleRouter.PUT("updateJMheathyTitles", JMTitleApi.UpdateJMheathyTitles)    // 更新标题表
	}
	{
		JMTitleRouterWithoutRecord.GET("findJMheathyTitles", JMTitleApi.FindJMheathyTitles)        // 根据ID获取标题表
		JMTitleRouterWithoutRecord.GET("getJMheathyTitlesList", JMTitleApi.GetJMheathyTitlesList)  // 获取标题表列表
	}
	{
	    JMTitleRouterWithoutAuth.GET("getJMheathyTitlesPublic", JMTitleApi.GetJMheathyTitlesPublic)  // 标题表开放接口
	}
}
