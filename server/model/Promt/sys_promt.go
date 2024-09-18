// 自动生成模板Promt
package Promt
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 提示词 结构体  Promt
type Promt struct {
    global.GVA_MODEL
    PromtData  datatypes.JSON `json:"promtData" form:"promtData" gorm:"index;column:promt_data;comment:提示词;size:255;" binding:"required"swaggertype:"array,object"`  //提示词 
    SystemUserId  *int `json:"systemUserId" form:"systemUserId" gorm:"index;column:system_user_id;comment:管理用户ID;size:8;" binding:"required"`  //管理用户ID 
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 提示词 Promt自定义表名 promt
func (Promt) TableName() string {
    return "promt"
}

