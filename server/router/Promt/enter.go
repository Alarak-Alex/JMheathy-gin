package Promt

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ PromtRouter }

var pmtApi = api.ApiGroupApp.PromtApiGroup.PromtApi
