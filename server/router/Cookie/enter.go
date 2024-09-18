package Cookie

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ CookieRouter }

var JMCookieApi = api.ApiGroupApp.CookieApiGroup.CookieApi
