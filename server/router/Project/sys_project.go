package Project

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SystemProjectRouter struct{}

func (s *SystemProjectRouter) InitSystemProjectRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	ProjectsRouter := Router.Group("Projects").Use(middleware.OperationRecord())
	ProjectsRouterWithoutRecord := Router.Group("Projects")
	ProjectsRouterWithoutAuth := PublicRouter.Group("Projects")
	{
		ProjectsRouter.POST("createSystemProject", ProjectsApi.CreateSystemProject)
		ProjectsRouter.DELETE("deleteSystemProject", ProjectsApi.DeleteSystemProject)
		ProjectsRouter.DELETE("deleteSystemProjectByIds", ProjectsApi.DeleteSystemProjectByIds)
		ProjectsRouter.PUT("updateSystemProject", ProjectsApi.UpdateSystemProject)
	}
	{
		ProjectsRouterWithoutRecord.GET("findSystemProject", ProjectsApi.FindSystemProject)
		ProjectsRouterWithoutRecord.GET("getSystemProjectList", ProjectsApi.GetSystemProjectList)
	}
	{
		ProjectsRouterWithoutAuth.GET("getSystemProjectPublic", ProjectsApi.GetSystemProjectPublic)
		ProjectsRouterWithoutAuth.POST("WriteWord", ProjectsApi.WriteWord)
		ProjectsRouterWithoutAuth.POST("PublishArticle", ProjectsApi.PublishArticle)
	}
}
