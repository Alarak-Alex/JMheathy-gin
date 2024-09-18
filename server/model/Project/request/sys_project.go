package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type SystemProjectSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    PicType  string `json:"picType" form:"picType" `
    PromtId  *int `json:"promtId" form:"promtId" `
    CookieType  string `json:"cookieType" form:"cookieType" `
    SystemUserId  *int `json:"systemUserId" form:"systemUserId" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
