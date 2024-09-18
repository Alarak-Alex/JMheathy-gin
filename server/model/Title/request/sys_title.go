package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type JMheathyTitlesSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    Title  string `json:"title" form:"title" `
    UploadUser  string `json:"uploadUser" form:"uploadUser" `
    UseTime  *int `json:"useTime" form:"useTime" `
    IsBan  string `json:"isBan" form:"isBan" `
    BanPlatform  string `json:"banPlatform" form:"banPlatform" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
