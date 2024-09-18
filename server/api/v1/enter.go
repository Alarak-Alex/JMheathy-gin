package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Cookie"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Picture"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Project"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Promt"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Title"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	CookieApiGroup  Cookie.ApiGroup
	PictureApiGroup Picture.ApiGroup
	PromtApiGroup   Promt.ApiGroup
	TitleApiGroup   Title.ApiGroup
	ProjectApiGroup Project.ApiGroup
}
