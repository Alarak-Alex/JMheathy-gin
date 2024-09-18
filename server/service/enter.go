package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/Cookie"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Picture"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Project"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Promt"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Title"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	CookieServiceGroup  Cookie.ServiceGroup
	PictureServiceGroup Picture.ServiceGroup
	PromtServiceGroup   Promt.ServiceGroup
	TitleServiceGroup   Title.ServiceGroup
	ProjectServiceGroup Project.ServiceGroup
}
