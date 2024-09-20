// 自动生成模板Picture
package Picture

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 图片 结构体  Picture
type Picture struct {
	global.GVA_MODEL
	Type       string `json:"type" form:"type" gorm:"index;default:1;column:type;comment:图片类型;size:255;" binding:"required"`         //图片类型
	Pic        string `json:"pic" form:"pic" gorm:"index;column:pic;comment:图片;size:255;" binding:"required"`                        //图片
	UpLoadUser uint   `json:"upLoadUser" form:"upLoadUser" gorm:"index;column:up_load_user;comment:上传用户;size:8;" binding:"required"` //上传用户
	CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 图片 Picture自定义表名 picture
func (Picture) TableName() string {
	return "picture"
}
