package Project

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ SystemProjectRouter }

var ProjectsApi = api.ApiGroupApp.ProjectApiGroup.SystemProjectApi
