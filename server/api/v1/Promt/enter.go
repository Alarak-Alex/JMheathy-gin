package Promt

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ PromtApi }

var pmtService = service.ServiceGroupApp.PromtServiceGroup.PromtService
