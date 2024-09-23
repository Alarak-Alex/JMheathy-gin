package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SystemProjectSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	PicType        string     `json:"picType" form:"picType" `
	PromtId        *int       `json:"promtId" form:"promtId" `
	CookieType     string     `json:"cookieType" form:"cookieType" `
	Status         string     `json:"status" form:"status" `
	SystemUserId   *int       `json:"systemUserId" form:"systemUserId" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
