// 自动生成模板SystemProject
package Project

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 项目 结构体  SystemProject
type SystemProject struct {
	global.GVA_MODEL
	TitleList    datatypes.JSON `json:"titleList" form:"titleList" gorm:"index;column:title_list;comment:标题列表;size:255;type:text;" binding:"required"swaggertype:"array,object"` //标题列表
	PicType      string         `json:"picType" form:"picType" gorm:"index;default:1;column:pic_type;comment:图片类型;size:255;" binding:"required"`                                 //图片类型
	PromtId      *int           `json:"promtId" form:"promtId" gorm:"index;column:promt_id;comment:提示词ID;size:8;" binding:"required"`                                            //提示词ID
	CookieType   string         `json:"cookieType" form:"cookieType" gorm:"index;column:cookie_type;comment:cookie类型;size:255;" binding:"required"`                              //cookie类型
	SystemUserId *int           `json:"systemUserId" form:"systemUserId" gorm:"index;column:system_user_id;comment:管理用户ID;size:8;" binding:"required"`                           //管理用户ID
	Status       string         `json:"status" form:"status" gorm:"index;default:5;column:status;comment:状态;size:8;" binding:"required"`                                         //状态
	PUUID        string         `json:"pUUID" form:"pUUID" gorm:"index;column:p_uuid;comment:项目UUID;size:255;"`                                                                  //项目UUID
	CreatedBy    uint           `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint           `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint           `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 项目 SystemProject自定义表名 system_project
func (SystemProject) TableName() string {
	return "system_project"
}
