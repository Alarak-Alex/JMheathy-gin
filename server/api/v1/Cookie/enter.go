package Cookie

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ CookieApi }

var JMCookieService = service.ServiceGroupApp.CookieServiceGroup.CookieService
