// 自动生成模板JMheathyTitles
package Title
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 标题表 结构体  JMheathyTitles
type JMheathyTitles struct {
    global.GVA_MODEL
    Title  string `json:"title" form:"title" gorm:"index;column:title;comment:标题;size:255;" binding:"required"`  //标题 
    UploadUser  string `json:"uploadUser" form:"uploadUser" gorm:"index;column:upload_user;comment:上传用户;size:255;" binding:"required"`  //上传用户 
    UseTime  *int `json:"useTime" form:"useTime" gorm:"index;default:1;column:use_time;comment:使用次数;size:8;"`  //使用次数 
    IsBan  string `json:"isBan" form:"isBan" gorm:"index;default:1;column:is_ban;comment:是否封禁;size:255;"`  //是否封禁 
    BanPlatform  string `json:"banPlatform" form:"banPlatform" gorm:"index;column:ban_platform;comment:封禁平台;size:255;" binding:"required"`  //封禁平台 
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 标题表 JMheathyTitles自定义表名 j_mheathy_titles
func (JMheathyTitles) TableName() string {
    return "j_mheathy_titles"
}

