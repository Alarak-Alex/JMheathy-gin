package Project

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ SystemProjectApi }

var ProjectsService = service.ServiceGroupApp.ProjectServiceGroup.SystemProjectService
