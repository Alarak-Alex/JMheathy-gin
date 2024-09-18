package Picture

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PictureRouter struct {}

// InitPictureRouter 初始化 图片 路由信息
func (s *PictureRouter) InitPictureRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	PicRouter := Router.Group("Pic").Use(middleware.OperationRecord())
	PicRouterWithoutRecord := Router.Group("Pic")
	PicRouterWithoutAuth := PublicRouter.Group("Pic")
	{
		PicRouter.POST("createPicture", PicApi.CreatePicture)   // 新建图片
		PicRouter.DELETE("deletePicture", PicApi.DeletePicture) // 删除图片
		PicRouter.DELETE("deletePictureByIds", PicApi.DeletePictureByIds) // 批量删除图片
		PicRouter.PUT("updatePicture", PicApi.UpdatePicture)    // 更新图片
	}
	{
		PicRouterWithoutRecord.GET("findPicture", PicApi.FindPicture)        // 根据ID获取图片
		PicRouterWithoutRecord.GET("getPictureList", PicApi.GetPictureList)  // 获取图片列表
	}
	{
	    PicRouterWithoutAuth.GET("getPicturePublic", PicApi.GetPicturePublic)  // 图片开放接口
	}
}
