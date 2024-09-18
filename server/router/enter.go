package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/Cookie"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Picture"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Project"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Promt"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Title"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Cookie  Cookie.RouterGroup
	Picture Picture.RouterGroup
	Promt   Promt.RouterGroup
	Title   Title.RouterGroup
	Project Project.RouterGroup
}
