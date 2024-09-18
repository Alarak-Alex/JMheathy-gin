package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CookieSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    CookieType  string `json:"cookieType" form:"cookieType" `
    SystemUserId  *int `json:"systemUserId" form:"systemUserId" `
    CookieName  string `json:"cookieName" form:"cookieName" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
