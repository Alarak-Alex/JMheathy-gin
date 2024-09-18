package Picture

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ PictureApi }

var PicService = service.ServiceGroupApp.PictureServiceGroup.PictureService
