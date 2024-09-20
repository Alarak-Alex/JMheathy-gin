package Picture

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PictureRouter struct{}

func (s *PictureRouter) InitPictureRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	PicRouter := Router.Group("Pic").Use(middleware.OperationRecord())
	PicRouterWithoutRecord := Router.Group("Pic")
	PicRouterWithoutAuth := PublicRouter.Group("Pic")
	{
		PicRouter.POST("createPicture", PicApi.CreatePicture)
		PicRouter.DELETE("deletePicture", PicApi.DeletePicture)
		PicRouter.DELETE("deletePictureByIds", PicApi.DeletePictureByIds)
		PicRouter.PUT("updatePicture", PicApi.UpdatePicture)
	}
	{
		PicRouterWithoutRecord.GET("findPicture", PicApi.FindPicture)
		PicRouterWithoutRecord.GET("getPictureList", PicApi.GetPictureList)
	}
	{
		PicRouterWithoutAuth.GET("getPicturePublic", PicApi.GetPicturePublic)
		PicRouterWithoutAuth.POST("CreateMorePic", PicApi.CreateMorePic)
	}
}
