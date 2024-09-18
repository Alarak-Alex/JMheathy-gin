package Promt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PromtRouter struct {}

// InitPromtRouter 初始化 提示词 路由信息
func (s *PromtRouter) InitPromtRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	pmtRouter := Router.Group("pmt").Use(middleware.OperationRecord())
	pmtRouterWithoutRecord := Router.Group("pmt")
	pmtRouterWithoutAuth := PublicRouter.Group("pmt")
	{
		pmtRouter.POST("createPromt", pmtApi.CreatePromt)   // 新建提示词
		pmtRouter.DELETE("deletePromt", pmtApi.DeletePromt) // 删除提示词
		pmtRouter.DELETE("deletePromtByIds", pmtApi.DeletePromtByIds) // 批量删除提示词
		pmtRouter.PUT("updatePromt", pmtApi.UpdatePromt)    // 更新提示词
	}
	{
		pmtRouterWithoutRecord.GET("findPromt", pmtApi.FindPromt)        // 根据ID获取提示词
		pmtRouterWithoutRecord.GET("getPromtList", pmtApi.GetPromtList)  // 获取提示词列表
	}
	{
	    pmtRouterWithoutAuth.GET("getPromtPublic", pmtApi.GetPromtPublic)  // 提示词开放接口
	}
}
