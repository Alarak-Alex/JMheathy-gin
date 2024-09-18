// 自动生成模板Cookie
package Cookie
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// Cookie 结构体  Cookie
type Cookie struct {
    global.GVA_MODEL
    UserCookie  datatypes.JSON `json:"userCookie" form:"userCookie" gorm:"index;column:user_cookie;comment:用户Cookie;size:255;" binding:"required"swaggertype:"array,object"`  //用户Cookie 
    CookieType  string `json:"cookieType" form:"cookieType" gorm:"index;column:cookie_type;comment:Cookie类型;size:255;" binding:"required"`  //Cookie类型 
    Enable  string `json:"enable" form:"enable" gorm:"index;default:1;column:enable;comment:是否启用;size:255;" binding:"required"`  //是否启用 
    SystemUserId  *int `json:"systemUserId" form:"systemUserId" gorm:"index;column:system_user_id;comment:管理用户ID;size:8;"`  //管理用户ID 
    CookieName  string `json:"cookieName" form:"cookieName" gorm:"index;column:cookie_name;comment:Cookie名;size:255;" binding:"required"`  //Cookie名 
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Cookie Cookie自定义表名 cookie
func (Cookie) TableName() string {
    return "cookie"
}

