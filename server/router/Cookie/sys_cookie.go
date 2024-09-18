package Cookie

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CookieRouter struct {}

// InitCookieRouter 初始化 Cookie 路由信息
func (s *CookieRouter) InitCookieRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	JMCookieRouter := Router.Group("JMCookie").Use(middleware.OperationRecord())
	JMCookieRouterWithoutRecord := Router.Group("JMCookie")
	JMCookieRouterWithoutAuth := PublicRouter.Group("JMCookie")
	{
		JMCookieRouter.POST("createCookie", JMCookieApi.CreateCookie)   // 新建Cookie
		JMCookieRouter.DELETE("deleteCookie", JMCookieApi.DeleteCookie) // 删除Cookie
		JMCookieRouter.DELETE("deleteCookieByIds", JMCookieApi.DeleteCookieByIds) // 批量删除Cookie
		JMCookieRouter.PUT("updateCookie", JMCookieApi.UpdateCookie)    // 更新Cookie
	}
	{
		JMCookieRouterWithoutRecord.GET("findCookie", JMCookieApi.FindCookie)        // 根据ID获取Cookie
		JMCookieRouterWithoutRecord.GET("getCookieList", JMCookieApi.GetCookieList)  // 获取Cookie列表
	}
	{
	    JMCookieRouterWithoutAuth.GET("getCookiePublic", JMCookieApi.GetCookiePublic)  // Cookie开放接口
	}
}
