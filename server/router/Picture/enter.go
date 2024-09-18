package Picture

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ PictureRouter }

var PicApi = api.ApiGroupApp.PictureApiGroup.PictureApi
